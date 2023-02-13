package appsync

import (
	"HomegrownDB/lib/datastructs"
	"sync"
)

type IdResolver[T datastructs.Number] struct {
	lock       sync.Locker
	idCounter  T
	missingIds []T
}

func NewIdResolver[T datastructs.Number](nextId T, missingIds []T) *IdResolver[T] {
	return &IdResolver[T]{
		lock:       &sync.Mutex{},
		idCounter:  nextId,
		missingIds: missingIds,
	}
}

func (ir *IdResolver[T]) NextId() T {
	ir.lock.Lock()
	defer ir.lock.Unlock()

	length := len(ir.missingIds)
	if length > 0 {
		nextId := ir.missingIds[length-1]
		ir.missingIds = ir.missingIds[:length-1]
		return nextId
	}

	id := ir.idCounter
	ir.idCounter++
	return id
}

func (ir *IdResolver[T]) AddId(missingId T) {
	ir.lock.Lock()
	defer ir.lock.Unlock()

	ir.missingIds = append(ir.missingIds, missingId)
}
