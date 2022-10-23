package buffer

import (
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"sync"
)

//todo implement intermediate buffer to which data from TableIO goes before it
// is saved in DBSharedBuffer, in this way bufferMapLock will be lock for shorter time
// see sharedBuffer.loadPage

func NewSharedBuffer(bufferSize uint, pageIOStore *pageio.Store) SharedBuffer {
	descriptorArray := make([]pageDescriptor, bufferSize)
	for i := uint(0); i < bufferSize; i++ {
		descriptorArray[i] = pageDescriptor{
			pageTag:          page.Tag{},
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
		bufferMap:     map[page.Tag]ArrayIndex{},
		bufferMapLock: &sync.RWMutex{},

		descriptorArray: descriptorArray,
		clock:           newClockSweep(descriptorArray),

		pageBufferArray: make([]byte, bufferSize*uint(page.Size)),

		ioStore: pageIOStore,
	}
}

type sharedBuffer struct {
	bufferMap     map[page.Tag]ArrayIndex
	bufferMapLock *sync.RWMutex

	descriptorArray []pageDescriptor
	clock           *clockSweep

	pageBufferArray []byte

	ioStore *pageio.Store
}

func (b *sharedBuffer) RPage(tag page.Tag) (page.TableRPage, error) {
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
	pageStart := uintptr(pageArrIndex) * uintptr(page.Size)
	return page.NewPage(nil, b.pageBufferArray[pageStart:pageStart+uintptr(page.Size)]), nil
}

func (b *sharedBuffer) WPage(tag page.Tag) (page.TableWPage, error) {
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
	pageStart := uintptr(pageArrIndex) * uintptr(page.Size)
	return page.NewPage(nil, b.pageBufferArray[pageStart:pageStart+uintptr(page.Size)]), nil
}

func (b *sharedBuffer) ReleaseWPage(tag page.Tag) {
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

func (b *sharedBuffer) ReleaseRPage(tag page.Tag) {
	b.bufferMapLock.RLock()
	index := b.bufferMap[tag]
	b.bufferMapLock.RUnlock()
	descriptor := &b.descriptorArray[index]

	descriptor.contentLock.RUnlock()
	descriptor.unpin()
}

// todo 1) razem z https://www.interdb.jp/pg/pgsql08.html#_8.4. 8.4.3 do chabra z pytaniami
// 2) prawdopodobnie zaimplementować własną hash mape
func (b *sharedBuffer) loadPage(tag page.Tag) (ArrayIndex, error) {
	for {
		victimIndex := b.clock.FindVictimPage()
		descriptor := &b.descriptorArray[victimIndex]

		pSize := uint(page.Size)
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

		err := b.ioStore.Get(tag.Relation).ReadPage(tag.PageId, arraySlot)
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

	err := b.ioStore.Get(descriptorTag.Relation).FlushPage(descriptorTag.PageId, pageData)
	if err != nil {
		return err
	}
	descriptor.descriptorLock.Lock()
	descriptor.isDirty = false
	descriptor.descriptorLock.Unlock()

	return nil
}

func handleFailedTableIO(tag page.Tag, arrayIndex ArrayIndex, err error) (ArrayIndex, error) {
	//todo implement me
	panic("Not implemented")
}
