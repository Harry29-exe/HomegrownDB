package queue

import "sync"

type Queue[T any] interface {
	Get() (T, bool)
	Push(elem T)
	PushAll(elems []T)
}

// BaseQueue is synchronized Queue implementation
// that does not guaranty any order of elements in it
type BaseQueue[T any] struct {
	elements    []T
	firstElemAt int
	lastElemAt  int
	lock        sync.Locker
}

func NewBaseQueue[T any](initialQueueSize int) *BaseQueue[T] {
	return &BaseQueue[T]{
		elements:    make([]T, initialQueueSize),
		firstElemAt: -1,
		lastElemAt:  -1,
		lock:        &sync.Mutex{},
	}
}

func (b *BaseQueue[T]) Get() (elem T, ok bool) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.firstElemAt < 0 {
		return elem, ok
	}

	elem = b.elements[b.firstElemAt]
	ok = true

	b.firstElemAt--
	if b.firstElemAt > b.lastElemAt {
		b.firstElemAt = -1
	}
	return
}

func (b *BaseQueue[T]) Push(elem T) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.firstElemAt > 0 {
		b.firstElemAt--
		b.elements[b.firstElemAt] = elem
	} else if b.lastElemAt+1 < len(b.elements) {
		if b.firstElemAt == -1 {
			b.firstElemAt = 0
		}
		b.lastElemAt++
		b.elements[b.lastElemAt] = elem
	} else {
		b.lastElemAt++
		b.elements = append(b.elements, elem)
	}
}

func (b *BaseQueue[T]) PushAll(elems []T) {
	b.lock.Lock()
	defer b.lock.Unlock()
	elemIndex := 0
	elemsLen := len(elems)

	for b.firstElemAt > 0 && elemIndex < elemsLen {
		b.firstElemAt--
		elemIndex++
		b.elements[b.firstElemAt] = elems[elemIndex]
	}
	if elemIndex == elemsLen {
		return
	}
	for b.lastElemAt < len(b.elements) && elemIndex < elemsLen {
		b.lastElemAt++
		elemIndex++
		b.elements[b.lastElemAt] = elems[elemIndex]
	}
	if elemIndex == elemsLen {
		return
	}
	b.elements = append(b.elements, elems[elemIndex:]...)
	b.lastElemAt += elemsLen - elemIndex
}
