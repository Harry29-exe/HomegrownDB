package buffer

import (
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"errors"
	"sync"
)

//todo implement intermediate buffer to which data from TableIO goes before it
// is saved in DBSharedBuffer, in this way bufferMapLock will be lock for shorter time
// see sharedBuff.loadPage

func newSharedBuffer(bufferSize uint, pageIOStore *pageio.Store) sharedBuffer {
	descriptorArray := make([]pageDescriptor, bufferSize)
	for i := uint(0); i < bufferSize; i++ {
		descriptorArray[i] = pageDescriptor{
			pageTag:          PageTag{},
			arrayIndex:       i,
			refCount:         0,
			usageCount:       0,
			isDirty:          false,
			contentLock:      sync.RWMutex{},
			ioInProgressLock: sync.Mutex{},
			descriptorLock:   0,
		}
	}

	return &sharedBuff{
		bufferMap:     map[PageTag]arrayIndex{},
		bufferMapLock: &sync.RWMutex{},

		descriptorArray: descriptorArray,
		clock:           newClockSweep(descriptorArray),

		pageBufferArray: make([]byte, bufferSize*uint(page.Size)),

		ioStore: pageIOStore,
	}
}

type sharedBuff struct {
	bufferMap     map[PageTag]arrayIndex
	bufferMapLock *sync.RWMutex

	descriptorArray []pageDescriptor
	clock           *clockSweep

	pageBufferArray []byte

	ioStore *pageio.Store
}

func (b *sharedBuff) RPage(tag PageTag) (buffPage, error) {
	b.bufferMapLock.RLock()

	pageArrIndex, ok := b.bufferMap[tag]
	var descriptor *pageDescriptor
	if ok {
		descriptor = &b.descriptorArray[pageArrIndex]
		descriptor.pin()
		b.bufferMapLock.RUnlock()

		descriptor.contentLock.RLock()
		pageStart := uintptr(pageArrIndex) * uintptr(page.Size)
		return buffPage{
			bytes: b.pageBufferArray[pageStart : pageStart+uintptr(page.Size)],
			isNew: false,
		}, nil

	} else {
		b.bufferMapLock.RUnlock()
		//todo add locks to new loadPage impl
		return b.loadRPage(tag)
	}
}

func (b *sharedBuff) WPage(tag PageTag) (buffPage, error) {
	b.bufferMapLock.RLock()

	pageArrIndex, ok := b.bufferMap[tag]
	var descriptor *pageDescriptor
	if ok {
		descriptor = &b.descriptorArray[pageArrIndex]
		descriptor.pin()
		b.bufferMapLock.RUnlock()

		descriptor.contentLock.Lock()
		pageStart := uintptr(pageArrIndex) * uintptr(page.Size)
		return buffPage{
			bytes: b.pageBufferArray[pageStart : pageStart+uintptr(page.Size)],
			isNew: false,
		}, nil

	} else {
		b.bufferMapLock.RUnlock()
		//todo add locks to new loadPage impl
		return b.loadWPage(tag)
	}
}

func (b *sharedBuff) ReleaseWPage(tag PageTag) {
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

func (b *sharedBuff) ReleaseRPage(tag PageTag) {
	b.bufferMapLock.RLock()
	index := b.bufferMap[tag]
	b.bufferMapLock.RUnlock()
	descriptor := &b.descriptorArray[index]

	descriptor.contentLock.RUnlock()
	descriptor.unpin()
}

func (b *sharedBuff) loadWPage(tag PageTag) (buffPage, error) {
	descriptor, err := b.loadPage(tag, true)
	pageIsNew := false
	if err != nil {
		if errors.Is(err, pageio.NoPageErrorType) {
			pageIsNew = true
		} else {
			return buffPage{}, err
		}
	}
	descriptor.contentLock.Lock()

	pageStart := uintptr(descriptor.arrayIndex) * uintptr(page.Size)
	return buffPage{
		bytes: b.pageBufferArray[pageStart : pageStart+uintptr(page.Size)],
		isNew: pageIsNew,
	}, nil
}

func (b *sharedBuff) loadRPage(tag PageTag) (buffPage, error) {
	descriptor, err := b.loadPage(tag, false)
	if err != nil {
		return buffPage{}, err
	}
	descriptor.contentLock.RLock()

	pageStart := uintptr(descriptor.arrayIndex) * uintptr(page.Size)
	return buffPage{
		bytes: b.pageBufferArray[pageStart : pageStart+uintptr(page.Size)],
		isNew: false,
	}, nil
}

// loadPage returns requested page, returned page is already pined to prevent
// unexpected deletions
//
// todo 1) razem z https://www.interdb.jp/pg/pgsql08.html#_8.4. 8.4.3 do chabra z pytaniami
// 2) prawdopodobnie zaimplementować własną hash mape
func (b *sharedBuff) loadPage(tag PageTag, wMode bool) (*pageDescriptor, error) {
	for {
		victimIndex := b.clock.FindVictimPage()
		descriptor := &b.descriptorArray[victimIndex]

		pSize := uint(page.Size)
		pageStart := pSize * victimIndex
		arraySlot := b.pageBufferArray[pageStart : pageStart+pSize]

		if descriptor.isDirty {
			err := b.flushPage(descriptor, arraySlot)
			if err != nil {
				return nil, err
			}
		}

		b.bufferMapLock.Lock()
		if descriptor.refCount != 0 {
			b.bufferMapLock.Unlock()
			continue
		}

		delete(b.bufferMap, descriptor.pageTag)
		b.bufferMap[tag] = victimIndex
		descriptor.Refresh(tag)
		descriptor.contentLock.Lock()
		descriptor.pin()
		//todo check if ioLock should not be locked here
		b.bufferMapLock.Unlock()

		err := b.ioStore.Get(tag.Relation).ReadPage(tag.PageId, arraySlot)
		descriptor.contentLock.Unlock()

		return descriptor, err
	}
}

func (b *sharedBuff) flushPage(descriptor *pageDescriptor, pageData []byte) error {
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

func handleFailedTableIO(tag PageTag, arrayIndex arrayIndex, err error) (arrayIndex, error) {
	//todo implement me
	panic("Not implemented")
}
