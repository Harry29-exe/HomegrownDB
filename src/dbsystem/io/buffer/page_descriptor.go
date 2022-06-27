package buffer

import (
	"HomegrownDB/datastructs/appsync"
	"HomegrownDB/dbsystem/bstructs"
	"sync"
)

type pageDescriptor struct {
	pageTag    bstructs.PageTag
	arrayIndex ArrayIndex
	refCount   uint32
	usageCount uint32

	contentLock      sync.RWMutex
	ioInProgressLock sync.Mutex
	descriptorLock   appsync.SpinLock
}

func (pd *pageDescriptor) InitNewPage(tag bstructs.PageTag) {
	pd.refCount = 0
	pd.usageCount = 0
	pd.pageTag = tag
}

// newUsage increment refCount and usageCount by 1
func (pd *pageDescriptor) newUsage() {
	pd.descriptorLock.Lock()
	pd.refCount++
	pd.usageCount++
	pd.descriptorLock.Unlock()
}

// endUsage decrement refCount by 1
func (pd *pageDescriptor) endUsage() {
	pd.descriptorLock.Lock()
	pd.refCount--
	pd.descriptorLock.Unlock()
}
