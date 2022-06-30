package buffer

import (
	"HomegrownDB/dbsystem/bstructs"
	"HomegrownDB/dbsystem/io"
	"HomegrownDB/dbsystem/schema"
	"HomegrownDB/dbsystem/schema/table"
	"sync"
)

func newBuffer(bufferSize uint) *buffer {
	descriptorArray := make([]pageDescriptor, bufferSize)

	return &buffer{
		bufferMap:     map[bstructs.PageTag]ArrayIndex{},
		bufferMapLock: &sync.RWMutex{},

		descriptorArray: descriptorArray,
		clock:           newClockSweep(descriptorArray),

		pageBufferArray: make([]byte, bufferSize*uint(bstructs.PageSize)),
	}
}

type buffer struct {
	bufferMap     map[bstructs.PageTag]ArrayIndex
	bufferMapLock *sync.RWMutex

	descriptorArray []pageDescriptor
	clock           *clockSweep

	pageBufferArray []byte
}

func (b *buffer) RPage(tag bstructs.PageTag) (bstructs.RPage, error) {
	tableDef := schema.Tables.Table(tag.TableId)
	b.bufferMapLock.RLock()

	pageArrIndex, ok := b.bufferMap[tag]
	var descriptor *pageDescriptor
	if ok {
		descriptor = &b.descriptorArray[pageArrIndex]
		descriptor.pin()
		b.bufferMapLock.RUnlock()

		descriptor.contentLock.RLock()
	} else {
		b.bufferMapLock.RUnlock()
		//todo add locks to new loadPage impl
		index, err := b.loadPage(tag, tableDef)
		if err != nil {
			return nil, err
		}

		descriptor = &b.descriptorArray[index]
	}

	pageStart := uintptr(pageArrIndex) * uintptr(bstructs.PageSize)
	return bstructs.NewPage(tableDef, b.pageBufferArray[pageStart:pageStart+uintptr(bstructs.PageSize)]), nil
}

func (b *buffer) WPage(tag bstructs.PageTag) (bstructs.WPage, error) {
	tableDef := schema.Tables.Table(tag.TableId)
	b.bufferMapLock.RLock()

	pageArrIndex, ok := b.bufferMap[tag]
	var descriptor *pageDescriptor
	if ok {
		descriptor = &b.descriptorArray[pageArrIndex]
		descriptor.pin()
		b.bufferMapLock.RUnlock()

		descriptor.contentLock.Lock()
	} else {
		b.bufferMapLock.RUnlock()
		//todo add locks to new loadPage impl
		index, err := b.loadPage(tag, tableDef)
		if err != nil {
			return nil, err
		}

		descriptor = &b.descriptorArray[index]
	}

	pageStart := uintptr(pageArrIndex) * uintptr(bstructs.PageSize)
	return bstructs.NewPage(tableDef, b.pageBufferArray[pageStart:pageStart+uintptr(bstructs.PageSize)]), nil
}

func (b *buffer) ReleaseWPage(page bstructs.WPage) {
	panic("Not implemented")
}

func (b *buffer) ReleaseRPage(page bstructs.RPage) {
	panic("Not implemented")
}

//todo 1) razem z https://www.interdb.jp/pg/pgsql08.html#_8.4. 8.4.3 do chabra z pytaniami
// 2) prawdopodobnie zaimplementować własną hash mape
func (b *buffer) loadPage(tag bstructs.PageTag, table table.Definition) (ArrayIndex, error) {
	for {
		victimIndex := b.clock.FindVictimPage()
		descriptor := &b.descriptorArray[victimIndex]

		pSize := uint(bstructs.PageSize)
		pageStart := pSize * victimIndex
		arraySlot := b.pageBufferArray[pageStart : pageStart+pSize]

		if descriptor.isDirty {
			descriptor.pin()
			descriptor.contentLock.RLock()
			descriptor.ioInProgressLock.Lock()

			io.Pages.Flush(descriptor.pageTag, arraySlot)
			descriptor.descriptorLock.Lock()
			descriptor.isDirty = false
			descriptor.descriptorLock.Unlock()

			descriptor.contentLock.RUnlock()
			descriptor.ioInProgressLock.Unlock()
			descriptor.unpin()
		}

		b.bufferMapLock.Lock()
		if descriptor.refCount != 0 {
			b.bufferMapLock.Unlock()
			continue
		}

		b.bufferMap[tag] = victimIndex
		delete(b.bufferMap, descriptor.pageTag)

		//todo load new page
	}
}
