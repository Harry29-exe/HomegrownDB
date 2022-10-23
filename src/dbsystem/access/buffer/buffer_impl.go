package buffer

import (
	"HomegrownDB/dbsystem/access"
	dbbs2 "HomegrownDB/dbsystem/access/dbbs"
	"HomegrownDB/dbsystem/schema/table"
	"sync"
)

//todo implement intermediate buffer to which data from TableIO goes before it
// is saved in SharedBuffer, in this way bufferMapLock will be lock for shorter time
// see sharedBuffer.loadPage

func NewSharedBuffer(bufferSize uint, tableStore table.Store, tableIO access.TableIOStore) DBSharedBuffer {
	descriptorArray := make([]pageDescriptor, bufferSize)
	for i := uint(0); i < bufferSize; i++ {
		descriptorArray[i] = pageDescriptor{
			pageTag:          dbbs2.PageTag{},
			arrayIndex:       i,
			refCount:         0,
			usageCount:       0,
			isDirty:          false,
			contentLock:      sync.RWMutex{},
			ioInProgressLock: sync.Mutex{},
			descriptorLock:   0,
		}
	}

	return &sharedBuffer{
		bufferMap:     map[dbbs2.PageTag]ArrayIndex{},
		bufferMapLock: &sync.RWMutex{},

		descriptorArray: descriptorArray,
		clock:           newClockSweep(descriptorArray),

		pageBufferArray: make([]byte, bufferSize*uint(dbbs2.PageSize)),

		tableStore:   tableStore,
		tableIOStore: tableIO,
	}
}

type sharedBuffer struct {
	bufferMap     map[dbbs2.PageTag]ArrayIndex
	bufferMapLock *sync.RWMutex

	descriptorArray []pageDescriptor
	clock           *clockSweep

	pageBufferArray []byte

	tableStore   table.Store
	tableIOStore access.TableIOStore
}

func (b *sharedBuffer) RPage(tag dbbs2.PageTag) (dbbs2.RPage, error) {
	tableDef := b.tableStore.Table(tag.Relation)
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

	descriptor.contentLock.RLock()
	descriptor.pin()
	pageStart := uintptr(pageArrIndex) * uintptr(dbbs2.PageSize)
	return dbbs2.NewPage(tableDef, b.pageBufferArray[pageStart:pageStart+uintptr(dbbs2.PageSize)]), nil
}

func (b *sharedBuffer) WPage(tag dbbs2.PageTag) (dbbs2.WPage, error) {
	tableDef := b.tableStore.Table(tag.Relation)
	b.bufferMapLock.RLock()

	pageArrIndex, ok := b.bufferMap[tag]
	var descriptor *pageDescriptor
	if ok {
		descriptor = &b.descriptorArray[pageArrIndex]
		b.bufferMapLock.RUnlock()
	} else {
		b.bufferMapLock.RUnlock()
		//todo add locks to new loadPage impl
		index, err := b.loadPage(tag)
		if err != nil {
			return nil, err
		}

		descriptor = &b.descriptorArray[index]
		pageArrIndex = index
	}

	descriptor.pin()
	descriptor.contentLock.Lock()
	pageStart := uintptr(pageArrIndex) * uintptr(dbbs2.PageSize)
	return dbbs2.NewPage(tableDef, b.pageBufferArray[pageStart:pageStart+uintptr(dbbs2.PageSize)]), nil
}

func (b *sharedBuffer) ReleaseWPage(tag dbbs2.PageTag) {
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

func (b *sharedBuffer) ReleaseRPage(tag dbbs2.PageTag) {
	b.bufferMapLock.RLock()
	index := b.bufferMap[tag]
	b.bufferMapLock.RUnlock()
	descriptor := &b.descriptorArray[index]

	descriptor.contentLock.RUnlock()
	descriptor.unpin()
}

// todo 1) razem z https://www.interdb.jp/pg/pgsql08.html#_8.4. 8.4.3 do chabra z pytaniami
// 2) prawdopodobnie zaimplementować własną hash mape
func (b *sharedBuffer) loadPage(tag dbbs2.PageTag) (ArrayIndex, error) {
	for {
		victimIndex := b.clock.FindVictimPage()
		descriptor := &b.descriptorArray[victimIndex]

		pSize := uint(dbbs2.PageSize)
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

		delete(b.bufferMap, descriptor.pageTag)
		b.bufferMap[tag] = victimIndex
		descriptor.InitNewPage(tag)
		descriptor.contentLock.Lock()
		//todo check if ioLock should not be locked here
		b.bufferMapLock.Unlock()

		err := b.tableIOStore.TableIO(tag.Relation).ReadPage(tag.PageId, arraySlot)
		if err != nil {
			return handleFailedTableIO(tag, victimIndex, err)
		}

		descriptor.contentLock.Unlock()
		return victimIndex, nil
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

	err := b.tableIOStore.TableIO(descriptorTag.Relation).FlushPage(descriptorTag.PageId, pageData)
	if err != nil {
		return err
	}
	descriptor.descriptorLock.Lock()
	descriptor.isDirty = false
	descriptor.descriptorLock.Unlock()

	return nil
}

func handleFailedTableIO(tag dbbs2.PageTag, arrayIndex ArrayIndex, err error) (ArrayIndex, error) {
	//todo implement me
	panic("Not implemented")
}
