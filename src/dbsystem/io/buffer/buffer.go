package buffer

import (
	"HomegrownDB/dbsystem/bstructs"
	"sync"
)

var Buffer = &buffer{
	pageIdBufferId:  map[bstructs.PageTag]ArrayIndex{},
	descriptorArray: make([]bufferDescriptor, 0, bufferSize),
	pageBufferArray: make([]byte, 0, int64(bstructs.PageSize)*bufferSize),
}

const bufferSize = 10_000

type ArrayIndex = uint

type bufferDescriptor struct {
	mutex sync.RWMutex
}
