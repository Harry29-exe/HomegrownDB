package appsync

import (
	"runtime"
	"sync/atomic"
)

type SpinLock uint32

func (sl *SpinLock) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		runtime.Gosched()
	}
}
func (sl *SpinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}
func NewSpinLock() *SpinLock {
	var lock SpinLock
	return &lock
}
