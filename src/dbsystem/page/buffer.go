package page

import (
	"sync"
)

var Buffer = &buffer{
	pageIdBufferId:  map[Tag]BufferId{},
	descriptorArray: make([]bufferDescriptor, 0, bufferSize),
	pageBufferArray: make([]byte, 0, Size*bufferSize),
}

const bufferSize = 10_000

type buffer struct {
	pageIdBufferId  map[Tag]BufferId
	descriptorArray []bufferDescriptor
	pageBufferArray []byte
}

type BufferId = uint64

type bufferDescriptor struct {
	mutex sync.RWMutex
}
