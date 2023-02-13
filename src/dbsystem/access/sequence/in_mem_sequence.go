package sequence

import (
	"HomegrownDB/lib/datastructs"
)

func NewInMemSequence[T datastructs.Number](nextVal T) Sequence[T] {
	return &InMemorySequence[T]{NextVal: nextVal}
}

type InMemorySequence[T datastructs.Number] struct {
	NextVal T
}

func (i *InMemorySequence[T]) Next() T {
	val := i.NextVal
	i.NextVal++
	return val
}

var _ Sequence[int] = &InMemorySequence[int]{}
