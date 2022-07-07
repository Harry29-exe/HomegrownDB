package buffer

import (
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/schema/table"
	"sync"
)

func NewSharedBuffer(bufferSize uint, tableSource TableSrc, pageLoader PageIO) *sharedBuffer {
	descriptorArray := make([]pageDescriptor, bufferSize)

	return &sharedBuffer{
		bufferMap:     map[bdata.PageTag]ArrayIndex{},
		bufferMapLock: &sync.RWMutex{},

		descriptorArray: descriptorArray,
		clock:           newClockSweep(descriptorArray),

		pageBufferArray: make([]byte, bufferSize*uint(bdata.PageSize)),

		tableSRC: tableSource,
		pageIO:   pageLoader,
	}
}

type sharedBuffer struct {
	bufferMap     map[bdata.PageTag]ArrayIndex
	bufferMapLock *sync.RWMutex

	descriptorArray []pageDescriptor
	clock           *clockSweep

	pageBufferArray []byte

	tableSRC TableSrc
	pageIO   PageIO
}

func (b *sharedBuffer) RPage(tag bdata.PageTag) (bdata.RPage, error) {
	tableDef := b.tableSRC.Table(tag.TableId)
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

	pageStart := uintptr(pageArrIndex) * uintptr(bdata.PageSize)
	return bdata.NewPage(tableDef, b.pageBufferArray[pageStart:pageStart+uintptr(bdata.PageSize)]), nil
}

func (b *sharedBuffer) WPage(tag bdata.PageTag) (bdata.WPage, error) {
	tableDef := b.tableSRC.Table(tag.TableId)
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

	pageStart := uintptr(pageArrIndex) * uintptr(bdata.PageSize)
	return bdata.NewPage(tableDef, b.pageBufferArray[pageStart:pageStart+uintptr(bdata.PageSize)]), nil
}

func (b *sharedBuffer) ReleaseWPage(tag bdata.PageTag) {
	b.bufferMapLock.RLock()
	index := b.bufferMap[tag]
	b.bufferMapLock.RUnlock()
	descriptor := &b.descriptorArray[index]

	descriptor.descriptorLock.Lock()
	descriptor.isDirty = true
	descriptor.descriptorLock.Unlock()

	descriptor.contentLock.Unlock()
	descriptor.unpin()
}

func (b *sharedBuffer) ReleaseRPage(tag bdata.PageTag) {
	b.bufferMapLock.RLock()
	index := b.bufferMap[tag]
	b.bufferMapLock.RUnlock()
	descriptor := &b.descriptorArray[index]

	descriptor.contentLock.RUnlock()
	descriptor.unpin()
}

//todo 1) razem z https://www.interdb.jp/pg/pgsql08.html#_8.4. 8.4.3 do chabra z pytaniami
// 2) prawdopodobnie zaimplementować własną hash mape
func (b *sharedBuffer) loadPage(tag bdata.PageTag, table table.Definition) (ArrayIndex, error) {
	for {
		victimIndex := b.clock.FindVictimPage()
		descriptor := &b.descriptorArray[victimIndex]

		pSize := uint(bdata.PageSize)
		pageStart := pSize * victimIndex
		arraySlot := b.pageBufferArray[pageStart : pageStart+pSize]

		if descriptor.isDirty {
			descriptor.pin()
			descriptor.contentLock.RLock()
			descriptor.ioInProgressLock.Lock()

			b.pageIO.Flush(descriptor.pageTag, arraySlot)
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
		b.pageIO.Read(tag, arraySlot)
		b.bufferMapLock.Unlock()
	}
}
