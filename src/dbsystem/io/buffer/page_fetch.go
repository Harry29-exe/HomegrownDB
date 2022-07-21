package buffer

import (
	"sync"
)

type clockSweep struct {
	lock            sync.Locker
	size            uint
	ptr             uint
	descriptorArray []pageDescriptor
}

func newClockSweep(descriptorArray []pageDescriptor) *clockSweep {
	return &clockSweep{
		lock:            &sync.Mutex{},
		size:            uint(len(descriptorArray)),
		ptr:             0,
		descriptorArray: descriptorArray,
	}
}

func (c clockSweep) FindVictimPage() (victimIndex uint) {
	var descriptor *pageDescriptor
	for {
		descriptor = &c.descriptorArray[c.ptr]
		descriptor.descriptorLock.Lock()
		if descriptor.usageCount == 0 {
			descriptor.descriptorLock.Unlock()

			victimIndex = c.ptr
			c.ptr = (c.ptr + 1) % c.size
			return
		} else {
			descriptor.usageCount--
			descriptor.descriptorLock.Unlock()
			c.ptr = (c.ptr + 1) % c.size
		}
	}
}
