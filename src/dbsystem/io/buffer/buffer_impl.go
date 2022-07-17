package buffer

import (
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/stores"
	"sync"
)

func NewSharedBuffer(bufferSize uint, tableStore stores.Tables) *sharedBuffer {
	descriptorArray := make([]pageDescriptor, bufferSize)

	return &sharedBuffer{
		bufferMap:     map[bdata.PageTag]ArrayIndex{},
		bufferMapLock: &sync.RWMutex{},

		descriptorArray: descriptorArray,
		clock:           newClockSweep(descriptorArray),

		pageBufferArray: make([]byte, bufferSize*uint(bdata.PageSize)),

		tableStore: tableStore,
	}
}

type sharedBuffer struct {
	bufferMap     map[bdata.PageTag]ArrayIndex
	bufferMapLock *sync.RWMutex

	descriptorArray []pageDescriptor
	clock           *clockSweep

	pageBufferArray []byte

	tableStore stores.Tables
}

func (b *sharedBuffer) RPage(tag bdata.PageTag) (bdata.RPage, error) {
	tableDef := b.tableStore.Table(tag.TableId)
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
		index, err := b.loadPage(tag)
		if err != nil {
			return nil, err
		}

		descriptor = &b.descriptorArray[index]
	}

	pageStart := uintptr(pageArrIndex) * uintptr(bdata.PageSize)
	return bdata.NewPage(tableDef, b.pageBufferArray[pageStart:pageStart+uintptr(bdata.PageSize)]), nil
}

func (b *sharedBuffer) WPage(tag bdata.PageTag) (bdata.WPage, error) {
	tableDef := b.tableStore.Table(tag.TableId)
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
		index, err := b.loadPage(tag)
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
func (b *sharedBuffer) loadPage(tag bdata.PageTag) (ArrayIndex, error) {
	for {
		victimIndex := b.clock.FindVictimPage()
		descriptor := &b.descriptorArray[victimIndex]

		pSize := uint(bdata.PageSize)
		pageStart := pSize * victimIndex
		arraySlot := b.pageBufferArray[pageStart : pageStart+pSize]

		if descriptor.isDirty {
			err := b.flushPage(descriptor, arraySlot)
			if err != nil {
				return 0, err
			}
		}

		b.bufferMapLock.Lock()
		if descriptor.refCount != 0 {
			b.bufferMapLock.Unlock()
			continue
		}

		b.bufferMap[tag] = victimIndex
		delete(b.bufferMap, descriptor.pageTag)

		//todo load new page
		err := b.tableStore.TableIO(tag.TableId).ReadPage(tag.PageId, arraySlot)
		if err != nil {
			b.bufferMapLock.Unlock()
			return 0, err
		}
		b.bufferMapLock.Unlock()
	}
}

func (b *sharedBuffer) flushPage(descriptor *pageDescriptor, pageData []byte) error {
	descriptorTag := descriptor.pageTag
	descriptor.pin()
	descriptor.contentLock.RLock()
	descriptor.ioInProgressLock.Lock()
	defer func() {
		descriptor.contentLock.RUnlock()
		descriptor.ioInProgressLock.Unlock()
		descriptor.unpin()
	}()

	err := b.tableStore.TableIO(descriptorTag.TableId).FlushPage(descriptorTag.PageId, pageData)
	if err != nil {
		return err
	}
	descriptor.descriptorLock.Lock()
	descriptor.isDirty = false
	descriptor.descriptorLock.Unlock()

	return nil
}
