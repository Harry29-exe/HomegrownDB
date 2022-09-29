package appsync

import (
	"HomegrownDB/common/datastructs"
	"sync"
	"sync/atomic"
)

type SyncCounter[T datastructs.Number] interface {
	GetAndIncrement() T
	IncrementAndGet() T
	Get() T
}

func NewSyncCounter[T datastructs.Number](start T) SyncCounter[T] {
	return &GenericSyncCounter[T]{
		val:  start,
		lock: 0,
	}
}

type GenericSyncCounter[T datastructs.Number] struct {
	val  T
	lock SpinLock
}

func (g *GenericSyncCounter[T]) GetAndIncrement() (v T) {
	g.lock.Lock()
	v = g.val
	g.val++
	g.lock.Unlock()
	return
}

func (g *GenericSyncCounter[T]) IncrementAndGet() (v T) {
	g.lock.Lock()
	g.val++
	v = g.val
	g.lock.Unlock()
	return
}

func (g *GenericSyncCounter[T]) Get() T {
	return g.val
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

func NewUint32SyncCounter(startVal uint32) SyncCounter[uint32] {
	return &uint32AtomicCounter{value: startVal}
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

type uint32AtomicCounter struct {
	value uint32
}

func (u *uint32AtomicCounter) GetAndIncrement() uint32 {
	return atomic.AddUint32(&u.value, 1) - 1
}

func (u *uint32AtomicCounter) IncrementAndGet() uint32 {
	return atomic.AddUint32(&u.value, 1)
}

func (u *uint32AtomicCounter) Get() uint32 {
	return u.value
}
