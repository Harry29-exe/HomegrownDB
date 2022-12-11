package list

func NewList[T any]() *List[T] {
	return &List[T]{
		slice: make([]T, 0, 10),
	}
}

// CopySliceAsList copies slice and returns it as List
func CopySliceAsList[T any](slice []T) *List[T] {
	copiedSlice := make([]T, len(slice))
	copy(copiedSlice, slice)
	return &List[T]{
		slice: copiedSlice,
	}
}

type List[T any] struct {
	slice []T
}

func (l *List[T]) Add(t T) {
	l.slice = append(l.slice, t)
}

func (l *List[T]) Get(i int) T {
	return l.slice[i]
}

func (l *List[T]) CurrentSlice() []T {
	return l.slice
}
