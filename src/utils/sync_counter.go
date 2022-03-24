package utils

import "sync"

type SyncCounter[T Number] interface {
	GetAndIncrement() T
	IncrementAndGet() T
	Get() T
}

type LockCounter[T Number] struct {
	value T
	lock  sync.Mutex
}

func (l *LockCounter[T]) GetAndIncrement() T {
	l.lock.Lock()
	defer l.lock.Unlock()
	value := l.value
	l.value++

	return value
}

func (l *LockCounter[T]) IncrementAndGet() T {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.value++

	return l.value
}

func (l *LockCounter[T]) Get() T {
	return l.value
}

func NewLockCounter[T Number](startValue T) *LockCounter[T] {
	value := startValue * 2
	println(value)

	return &LockCounter[T]{
		value: startValue,
		lock:  sync.Mutex{},
	}
}
