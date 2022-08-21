package data

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/schema/column"
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

	var dataPosition = holder.Fields()*2 + 2
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
				field++
			}
		}
	}
	bparse.Serialize.PutUInt2(dataPosition, slot[field*2:])

	return Row{
		holder: holder,
		data:   slot,
	}
}

func (d Row) GetField(fieldIndex uint16) []byte {
	start := bparse.Parse.Int2(d.data[fieldIndex*2:])
	if start == 0 {
		return nil
	}
	end := bparse.Parse.Int2(d.data[fieldIndex*2+2:])
	return d.data[start:end]
}
