package data

import "HomegrownDB/datastructs/queue"

const initialBufferSize = 1000

var GlobalSlotBuffer = NewSlotBuffer(initialBufferSize)

func NewSlotBuffer(arrayLen int) *SlotBuffer {
	array := make([]byte, initialBufferSize*arrayLen)
	baseQueue := queue.NewBaseQueue[[]byte](initialBufferSize)
	for i := 0; i < initialBufferSize; i++ {
		baseQueue.Push(array[i*arrayLen : (i+1)*arrayLen])
	}

	return &SlotBuffer{
		arrayQueue: baseQueue,
		arrayLen:   arrayLen,
	}
}

type SlotBuffer struct {
	arrayQueue queue.Queue[[]byte]
	arrayLen   int
}

func (b *SlotBuffer) GetInfo() *RowBuffer {
	//todo implement me
	panic("Not implemented")
}

func (b *SlotBuffer) GetArray() []byte {
	arr, ok := b.arrayQueue.Get()
	if ok {
		return arr
	} else {
		b.allocateNewHugeArray()
		return b.GetArray()
	}
}

func (b *SlotBuffer) ArrayLen() int {
	return b.arrayLen
}

func (b *SlotBuffer) allocateNewHugeArray() {
	array := make([]byte, initialBufferSize*b.arrayLen)
	for i := 0; i < initialBufferSize; i++ {
		b.arrayQueue.Push(array[i*b.arrayLen : (i+1)*b.arrayLen])
	}
}
