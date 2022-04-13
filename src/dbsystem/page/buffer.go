package page

import (
	"encoding/binary"
	"encoding/gob"
	"sync"
)

var Buffer = &buffer{map[PageId]uint64{}}

const bufferSize = 10_000

type buffer struct {
	pageIdBufferId  map[PageId]BufferId
	descriptorArray []bufferDescriptor
	pageBufferArray[]
}

type BufferId = uint64

type bufferDescriptor struct {
	mutex sync.RWMutex
}

func we() {
	binary.LittleEndian.Uint16()
	binary.BigEndian.Uint64()
	gob.Encoder{}
}
