package collection

type MutCollection[T any] interface {
	Collection[T]
	Add(value T) bool
	AddAll(value Collection[T]) bool
	AddSlice(value []T) bool
	Remove(value T) bool
	RemoveAll(value Collection[T]) bool
	RemoveSlice(value []T) bool
}
