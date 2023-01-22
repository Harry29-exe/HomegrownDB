package buffer

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/storage/page"
	"sync"
)

type pageDescriptor struct {
	pageTag    page.PageTag
	slotIndex  slotIndex
	refCount   uint32
	usageCount uint32

	isDirty bool

	contentLock      sync.RWMutex
	ioInProgressLock sync.Mutex
	descriptorLock   appsync.SpinLock
}

func (pd *pageDescriptor) Refresh(tag page.PageTag) {
	pd.descriptorLock.Lock()
	pd.usageCount = 2 // set usageCount to 2, so it won't instantly become victim page
	pd.pageTag = tag
	pd.descriptorLock.Unlock()
}

// pin increment refCount and usageCount by 1
func (pd *pageDescriptor) pin() {
	pd.descriptorLock.Lock()
	pd.refCount++
	pd.usageCount++
	pd.descriptorLock.Unlock()
}

// unpin decrement refCount by 1
func (pd *pageDescriptor) unpin() {
	pd.descriptorLock.Lock()
	pd.refCount--
	pd.descriptorLock.Unlock()
}
