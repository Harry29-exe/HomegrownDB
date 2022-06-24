package appsync

import (
	"HomegrownDB/datastructs"
	"sync"
)

type SyncCounter[T datastructs.Number] interface {
	GetAndIncrement() T
	IncrementAndGet() T
	Get() T
}

func NewUint64SyncCounter(startVal uint64) SyncCounter[uint64] {
	return newUint64Counter(startVal)
}

func NewInt32SyncCounter(startVal int32) SyncCounter[int32] {
	return &int32LockCounter{
		mutex: sync.Mutex{},
		value: startVal,
	}
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

type int32LockCounter struct {
	mutex sync.Mutex
	value int32
}

func (u *int32LockCounter) GetAndIncrement() int32 {
	value := u.value
	u.mutex.Lock()
	u.value++
	u.mutex.Unlock()
	return value
}

func (u *int32LockCounter) IncrementAndGet() int32 {
	u.mutex.Lock()
	u.value++
	u.mutex.Unlock()
	return u.value
}

func (u *int32LockCounter) Get() int32 {
	return u.value
}
