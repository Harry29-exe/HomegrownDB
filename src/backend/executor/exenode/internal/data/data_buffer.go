package data

type Buffer struct {
	dataInfoSlots []RowHolder
}

func (b *Buffer) GetInfo() *RowHolder {
	//todo implement me
	panic("Not implemented")
}

func (b *Buffer) GetArray() []byte {
	//todo implement me
	panic("Not implemented")
}

func (b *Buffer) ArrayLen() int {
	//todo implement me
	panic("Not implemented")
}
