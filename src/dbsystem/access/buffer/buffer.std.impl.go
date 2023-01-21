package buffer

import (
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"errors"
	"sync"
)

//todo implement intermediate buffer to which data from TableIO goes before it
// is saved in DBSharedBuffer, in this way bufferMapLock will be lock for shorter time
// see buffer.loadPage

func NewStdBuffer(bufferSize uint, pageIOStore pageio.Store) StdBuffer {
	descriptorArray := make([]pageDescriptor, bufferSize)
	for i := uint(0); i < bufferSize; i++ {
		descriptorArray[i] = pageDescriptor{
			pageTag:          pageio.PageTag{OwnerID: 0, PageId: page.InvalidId},
			slotIndex:        i,
			refCount:         0,
			usageCount:       0,
			isDirty:          false,
			contentLock:      sync.RWMutex{},
			ioInProgressLock: sync.Mutex{},
			descriptorLock:   0,
		}
	}

	return &buffer{
		bufferMap:     map[pageio.PageTag]slotIndex{},
		bufferMapLock: &sync.RWMutex{},

		descriptorArray: descriptorArray,
		clock:           newClockSweep(descriptorArray),

		pageBufferArray: make([]byte, bufferSize*uint(page.Size)),

		ioStore: pageIOStore,
	}
}

type buffer struct {
	bufferMap     map[pageio.PageTag]slotIndex
	bufferMapLock *sync.RWMutex

	descriptorArray []pageDescriptor
	clock           *clockSweep

	pageBufferArray []byte

	ioStore pageio.Store
}

func (b *buffer) ReadRPage(ownerID dbobj.OID, pageId page.Id, strategy rbm) (stdPage, error) {
	tag := pageio.PageTag{PageId: pageId, OwnerID: ownerID}
	b.bufferMapLock.RLock()

	pageArrIndex, ok := b.bufferMap[tag]
	if ok {
		descriptor := &b.descriptorArray[pageArrIndex]
		descriptor.pin()
		b.bufferMapLock.RUnlock()

		descriptor.contentLock.RLock()
		pageStart := uintptr(pageArrIndex) * uintptr(page.Size)
		return stdPage{
			Bytes: b.pageBufferArray[pageStart : pageStart+uintptr(page.Size)],
			IsNew: false,
		}, nil

	} else {
		b.bufferMapLock.RUnlock()
		return b.loadRPage(ownerID, pageId, strategy)
	}
}

func (b *buffer) ReadWPage(ownerID dbobj.OID, pageId page.Id, strategy rbm) (stdPage, error) {
	tag := pageio.PageTag{PageId: pageId, OwnerID: ownerID}
	b.bufferMapLock.RLock()

	pageArrIndex, ok := b.bufferMap[tag]
	var descriptor *pageDescriptor
	if ok {
		descriptor = &b.descriptorArray[pageArrIndex]
		descriptor.pin()
		b.bufferMapLock.RUnlock()

		descriptor.contentLock.Lock()
		pageStart := uintptr(pageArrIndex) * uintptr(page.Size)
		return stdPage{
			Bytes: b.pageBufferArray[pageStart : pageStart+uintptr(page.Size)],
			IsNew: false,
		}, nil

	} else {
		b.bufferMapLock.RUnlock()
		//todo add locks to new loadPage impl
		return b.loadWPage(ownerID, pageId, strategy)
	}
}

func (b *buffer) ReleaseWPage(tag pageio.PageTag) {
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

func (b *buffer) ReleaseRPage(tag pageio.PageTag) {
	b.bufferMapLock.RLock()
	index := b.bufferMap[tag]
	b.bufferMapLock.RUnlock()
	descriptor := &b.descriptorArray[index]

	descriptor.contentLock.RUnlock()
	descriptor.unpin()
}

func (b *buffer) loadWPage(ownerID dbobj.OID, pageId page.Id, strategy rbm) (stdPage, error) {
	descriptor, err := b.loadPage(ownerID, pageId)
	pageIsNew := false
	if err != nil {
		if errors.Is(err, pageio.NoPageErrorType) {
			pageIsNew = true
		} else {
			return stdPage{}, err
		}
	}
	descriptor.contentLock.Lock()

	pageStart := uintptr(descriptor.slotIndex) * uintptr(page.Size)
	return stdPage{
		Bytes: b.pageBufferArray[pageStart : pageStart+uintptr(page.Size)],
		IsNew: pageIsNew,
	}, nil
}

func (b *buffer) loadRPage(ownerID dbobj.OID, pageId page.Id, strategy rbm) (stdPage, error) {
	descriptor, err := b.loadPage(ownerID, pageId)
	descriptor.contentLock.RLock()
	if err != nil {
		println("pageId: ", pageId, ", ", err.Error())
		return stdPage{}, err
	}

	pageStart := uintptr(descriptor.slotIndex) * uintptr(page.Size)
	return stdPage{
		Bytes: b.pageBufferArray[pageStart : pageStart+uintptr(page.Size)],
		IsNew: false,
	}, nil
}

// loadPage returns requested page, returned page is already pined to prevent
// unexpected deletions
//
// todo 1) razem z https://www.interdb.jp/pg/pgsql08.html#_8.4. 8.4.3 do chabra z pytaniami
// 2) prawdopodobnie zaimplementować własną hash mape
func (b *buffer) loadPage(ownerID dbobj.OID, pageId page.Id) (*pageDescriptor, error) {
	pageTag := pageio.PageTag{OwnerID: ownerID, PageId: pageId}

	for {
		descriptor, err := b.prepareVictimPage()
		if err != nil {
			return nil, err
		}

		b.bufferMapLock.Lock()
		if pageId != NewPage {
			// checking if other goroutine didn't loaded page
			if arrIndex, ok := b.bufferMap[pageTag]; ok {
				descriptor.unpin()
				descriptor = &b.descriptorArray[arrIndex]
				descriptor.pin()
				b.bufferMapLock.Unlock()
				return descriptor, nil
			}

			// checking if other goroutine didn't start using this page
		} else if descriptor.refCount != 1 {
			descriptor.unpin()
			b.bufferMapLock.Unlock()
			continue
		}

		delete(b.bufferMap, descriptor.pageTag)
		relationIO := b.ioStore.Get(ownerID)
		if pageId == NewPage {
			pageTag.PageId = relationIO.PrepareNewPage()
		}
		b.bufferMap[pageTag] = descriptor.slotIndex
		descriptor.Refresh(pageTag)
		descriptor.contentLock.Lock()
		//todo check if ioLock should not be locked here

		b.bufferMapLock.Unlock()

		if pageId == NewPage {
			b.clearSlot(descriptor.slotIndex)
		} else {
			err = relationIO.ReadPage(pageId, b.getArraySlot(descriptor.slotIndex))
		}
		descriptor.contentLock.Unlock()

		return descriptor, err
	}
}

func (b *buffer) prepareVictimPage() (*pageDescriptor, error) {
	b.bufferMapLock.RLock()
	victimIndex := b.clock.FindVictimPage()
	b.bufferMapLock.RUnlock()
	descriptor := &b.descriptorArray[victimIndex]

	arraySlot := b.getArraySlot(victimIndex)

	if descriptor.isDirty {
		err := b.flushPage(descriptor, arraySlot)
		if err != nil {
			return nil, err
		}
	}
	return descriptor, nil
}

func (b *buffer) getArraySlot(slotIndex slotIndex) []byte {
	pSize := uint(page.Size)
	pageStart := pSize * slotIndex
	return b.pageBufferArray[pageStart : pageStart+pSize]
}

func (b *buffer) clearSlot(slotIndex slotIndex) {
	pageSlot := b.getArraySlot(slotIndex)
	for i := 0; i < int(page.Size); i++ {
		pageSlot[i] = 0
	}
}

func (b *buffer) flushPage(descriptor *pageDescriptor, pageData []byte) error {
	descriptorTag := descriptor.pageTag
	descriptor.pin()
	descriptor.contentLock.RLock()
	descriptor.ioInProgressLock.Lock()
	defer func() {
		descriptor.contentLock.RUnlock()
		descriptor.ioInProgressLock.Unlock()
		descriptor.unpin()
	}()

	err := b.ioStore.Get(descriptorTag.OwnerID).FlushPage(descriptorTag.PageId, pageData)
	if err != nil {
		return err
	}
	descriptor.descriptorLock.Lock()
	descriptor.isDirty = false
	descriptor.descriptorLock.Unlock()

	return nil
}
