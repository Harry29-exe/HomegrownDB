package page

import (
	"sync"
)

var Buffer = &buffer{
	pageIdBufferId:  map[Tag]BufferId{},
	descriptorArray: make([]bufferDescriptor, 0, bufferSize),
	pageBufferArray: make([]byte, 0, int64(Size)*bufferSize),
}

const bufferSize = 10_000

type BufferId = uint64

type bufferDescriptor struct {
	mutex sync.RWMutex
}
