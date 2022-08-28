package collection

type Collection[T any] interface {
	Get(object T) T
	AsSlice() []T

	Contain(object T) bool
	ContainF(object T, equals Equals[T]) bool
	ContainAll(objects Collection[T]) bool
	ContainAllF(objects Collection[T], equals Equals[T]) bool
	ContainSlice(objects []T) bool
	ContainSliceF(object []T, equals Equals[T]) bool

	Size() uint
}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type Equals[T any] func(val1, val2 T) bool
