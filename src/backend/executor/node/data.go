package node

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
)

type DataBuffer struct {
	dataInfoSlots []DataInfo
}

func (b *DataBuffer) GetInfo() *DataInfo {

}

func (b DataBuffer) GetArray() []byte {

}

type DataInfo struct {
	dataBuffer *DataBuffer

	tables    []table.Definition
	parsers   []column.DataParser
	fields    int16
	headerLen int

	dataArrays     [][]byte
	lastArrayIndex int
	lastArrayLen   int
}

func (i *DataInfo) GetDataSlot(dataContentLen int) []byte {
	dataLen := dataContentLen + i.headerLen
	freeSpace := cap(i.dataArrays[i.lastArrayIndex]) - i.lastArrayLen
	if freeSpace >= dataLen {
		data := i.dataArrays[i.lastArrayIndex][i.lastArrayLen : i.lastArrayLen+dataLen]
		i.lastArrayLen += dataLen
		return data
	}

	i.dataBuffer
}

type Data struct {
	data []byte // data format described in Data.svg
}

type DataFieldPtr = uint16

const DataFieldPtrSize = 2

func (d Data) GetField(fieldIndex int) []byte {

}
