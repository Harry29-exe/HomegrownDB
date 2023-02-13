package appsync

import "HomegrownDB/lib/datastructs"

func NewSimpleCounter[T datastructs.Number](next T) SimpleSyncCounter[T] {
	return SimpleSyncCounter[T]{
		val:  next,
		lock: 0,
	}
}

type SimpleSyncCounter[T datastructs.Number] struct {
	val  T
	lock SpinLock
}

func (s *SimpleSyncCounter[T]) Next() T {
	s.lock.Lock()
	next := s.val
	s.val++
	s.lock.Unlock()

	return next
}
