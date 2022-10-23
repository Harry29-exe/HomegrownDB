package qrow

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/access/dbbs"
	"HomegrownDB/dbsystem/schema/column"
)

type Row struct {
	holder RowBuffer
	data   []byte // data format described in documentation/executor/Data.svg
}

func NewRow(tuples []dbbs.Tuple, holder RowBuffer) Row {
	dataSize := 0
	for _, tuple := range tuples {
		dataSize = tuple.DataSize()
	}
	slot := holder.GetRowSlot(dataSize)

	var fieldCount = holder.Fields()
	var dataStart = fieldCount*2 + 2
	var dataPtr = dataStart
	var tuple dbbs.Tuple
	var val []byte
	field := FieldPtr(0)
	for i, table := range holder.Tables() {
		tuple = tuples[i]
		tupleData := tuple.Bytes()
		for colOrder, col := range table.Columns() {
			if tuple.IsNull(column.Order(colOrder)) {
				bparse.Serialize.PutUInt2(dataStart, slot[field*2:])
			} else {
				val, tupleData = col.CType().ValueAndSkip(tupleData)
				copy(slot[dataPtr:], val)
				bparse.Serialize.PutUInt2(dataPtr, slot[field*2:])
				dataPtr += uint16(len(val))
			}
			field++
		}
	}
	bparse.Serialize.PutUInt2(dataPtr, slot[field*2:])

	dataStartBytes := bparse.Serialize.Uint2(dataStart)
	for i := (fieldCount - 1) * 2; i > 0; i -= 2 {
		if slot[i] == dataStartBytes[0] && slot[i+1] == dataStartBytes[1] {
			slot[i] = slot[i+2]
			slot[i+1] = slot[i+3]
		}
	}

	return Row{
		holder: holder,
		data:   slot,
	}
}

func (d Row) GetField(fieldIndex uint16) []byte {
	start := bparse.Parse.Int2(d.data[fieldIndex*2:])
	end := bparse.Parse.Int2(d.data[fieldIndex*2+2:])
	if start == end {
		return nil
	}
	return d.data[start:end]
}
