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
