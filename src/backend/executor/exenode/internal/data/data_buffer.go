package data

import "HomegrownDB/datastructs/queue"

const initialBufferSize = 1000

func NewBuffer(arrayLen int) *Buffer {
	array := make([]byte, initialBufferSize*arrayLen)
	baseQueue := queue.NewBaseQueue[[]byte](initialBufferSize)
	for i := 0; i < initialBufferSize; i++ {
		baseQueue.Push(array[i*arrayLen : (i+1)*arrayLen])
	}

	return &Buffer{
		arrayQueue: baseQueue,
		arrayLen:   arrayLen,
	}
}

type Buffer struct {
	arrayQueue queue.Queue[[]byte]
	arrayLen   int
}

func (b *Buffer) GetInfo() *RowHolder {
	//todo implement me
	panic("Not implemented")
}

func (b *Buffer) GetArray() []byte {
	arr, ok := b.arrayQueue.Get()
	if ok {
		return arr
	} else {
		b.allocateNewHugeArray()
		return b.GetArray()
	}
}

func (b *Buffer) ArrayLen() int {
	return b.arrayLen
}

func (b *Buffer) allocateNewHugeArray() {
	array := make([]byte, initialBufferSize*b.arrayLen)
	for i := 0; i < initialBufferSize; i++ {
		b.arrayQueue.Push(array[i*b.arrayLen : (i+1)*b.arrayLen])
	}
}
