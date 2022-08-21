package data

import (
	"HomegrownDB/dbsystem/schema/table"
)

type RowHolder interface {
	GetRowSlot(dataContentLen int) []byte
	Free()
	Fields() uint16
	Tables() []table.Definition
}

type FieldPtr = uint16

const FieldPtrSize = 2

func NewBaseHolder(buffer *Buffer, tables []table.Definition) *BaseHolder {
	fields := uint16(0)
	for _, def := range tables {
		fields += def.ColumnCount()
	}

	return &BaseHolder{
		buffer:         buffer,
		tables:         tables,
		fields:         fields,
		headerLen:      int((fields + 1) * FieldPtrSize),
		dataArrays:     make([][]byte, 10),
		lastArrayIndex: 0,
		lastArrayLen:   0,
	}
}

type BaseHolder struct {
	buffer *Buffer

	tables    []table.Definition
	fields    uint16
	headerLen int

	dataArrays     [][]byte
	lastArrayIndex int
	lastArrayLen   int
}

func (i *BaseHolder) GetRowSlot(dataContentLen int) []byte {
	dataLen := dataContentLen + i.headerLen
	if dataLen > i.buffer.ArrayLen() {
		panic("Data row bigger thant array not yet supported")
	}

	freeSpace := cap(i.dataArrays[i.lastArrayIndex]) - i.lastArrayLen
	if freeSpace >= dataLen {
		data := i.dataArrays[i.lastArrayIndex][i.lastArrayLen : i.lastArrayLen+dataLen]
		i.lastArrayLen += dataLen
		return data
	}

	i.dataArrays = append(i.dataArrays, i.buffer.GetArray())
	i.lastArrayIndex = len(i.dataArrays) - 1
	i.lastArrayLen = dataLen
	return i.dataArrays[i.lastArrayIndex][0:dataLen]
}

func (i *BaseHolder) Free() {
	//todo implement me
	panic("Not implemented")
}

func (i *BaseHolder) Fields() uint16 {
	return i.fields
}

func (i *BaseHolder) Tables() []table.Definition {
	return i.tables
}
