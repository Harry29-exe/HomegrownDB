package qrow

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/schema/column"
)

type Row struct {
	holder RowBuffer
	data   []byte // data format described in documentation/executor/Data.svg
}

func NewRow(tuples []bdata.Tuple, holder RowBuffer) Row {
	dataSize := 0
	for _, tuple := range tuples {
		dataSize = tuple.DataSize()
	}
	slot := holder.GetRowSlot(dataSize)

	var fieldCount = holder.Fields()
	var dataPosition = fieldCount*2 + 2
	var tuple bdata.Tuple
	var tupleByte int
	field, skippedBytes := FieldPtr(0), 0
	for i, table := range holder.Tables() {
		tuple = tuples[i]
		for colOrder, parser := range table.AllColumnParsers() {
			if tuple.IsNull(column.OrderId(colOrder)) {
				slot[field*2] = 0
				slot[field*2+1] = 0
			} else {
				skippedBytes = parser.CopyData(tuple.Data()[tupleByte:], slot[dataPosition:])
				bparse.Serialize.PutUInt2(dataPosition, slot[field*2:])
				dataPosition += uint16(skippedBytes)
			}
			field++
		}
	}
	bparse.Serialize.PutUInt2(dataPosition, slot[field*2:])

	for i := (fieldCount - 1) * 2; i > 0; i -= 2 {
		if slot[i] == 0 && slot[i+1] == 0 {
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
