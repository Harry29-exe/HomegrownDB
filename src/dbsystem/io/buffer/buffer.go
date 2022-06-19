package buffer

import (
	"HomegrownDB/dbsystem/bstructs"
	"sync"
)

var Buffer = &buffer{
	pageIdBufferId:  map[bstructs.PageTag]BufferId{},
	descriptorArray: make([]bufferDescriptor, 0, bufferSize),
	pageBufferArray: make([]byte, 0, int64(bstructs.PageSize)*bufferSize),
}

const bufferSize = 10_000

type BufferId = uint64

type bufferDescriptor struct {
	mutex sync.RWMutex
}
