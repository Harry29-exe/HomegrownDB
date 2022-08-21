package data

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/bdata"
)

type Row struct {
	holder RowHolder
	data   []byte // data format described in documentation/executor/Data.svg
}

func NewRow(tuples []bdata.Tuple, holder RowHolder) Row {
	dataSize := 0
	for _, tuple := range tuples {
		dataSize = tuple.DataSize()
	}
	slot := holder.GetRowSlot(dataSize)

	var dataPosition = holder.Fields()
	var tuple []byte
	var tupleByte int
	field, skippedBytes := FieldPtr(0), 0
	for i, table := range holder.Tables() {
		tuple = tuples[i].Data()
		for _, parser := range table.AllColumnParsers() {
			skippedBytes = parser.CopyData(tuple[tupleByte:], slot[dataPosition:])
			dataPosition += uint16(skippedBytes)
			bparse.Serialize.PutUInt2(dataPosition, tuple[field*2:])
		}
	}

	return Row{
		holder: holder,
		data:   slot,
	}
}

func (d Row) GetField(fieldIndex int) []byte {
	start := bparse.Parse.Int2(d.data[fieldIndex*2:])
	end := bparse.Parse.Int2(d.data[fieldIndex*2+2:])
	return d.data[start:end]
}
