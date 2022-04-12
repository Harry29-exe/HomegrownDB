package datastructs

import "sync"

type SyncCounter[T Number] interface {
	GetAndIncrement() T
	IncrementAndGet() T
	Get() T
}

func NewUint64SyncCounter(startVal uint64) SyncCounter[uint64] {
	return newUint64Counter(startVal)
}

func newUint64Counter(startVal uint64) *uint64LockCounter {
	return &uint64LockCounter{
		mutex: sync.Mutex{},
		value: startVal,
	}
}

type uint64LockCounter struct {
	mutex sync.Mutex
	value uint64
}

func (u *uint64LockCounter) GetAndIncrement() uint64 {
	value := u.value
	u.mutex.Lock()
	u.value++
	u.mutex.Unlock()
	return value
}

func (u *uint64LockCounter) IncrementAndGet() uint64 {
	u.mutex.Lock()
	u.value++
	u.mutex.Unlock()
	return u.value
}

func (u *uint64LockCounter) Get() uint64 {
	return u.value
}
