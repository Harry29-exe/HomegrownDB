package dbbs

import (
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
	"math"
)

// QRow - query row is representation of data (often converted from tuple or
// concatenation of tuples) during data processing phase
type QRow struct {
	ptrs    []uint32
	data    []byte
	Pattern []ctype.CType
}

func NewQRowFromTuple(tuple RTuple) QRow {
	t := tuple.(Tuple)
	columns := t.table.Columns()

	headerLen := 4 * int(t.table.ColumnCount()+1)
	qrow := QRow{
		data:    make([]byte, t.DataSize()),
		ptrs:    make([]uint32, headerLen),
		Pattern: t.table.CTypePattern(),
	}

	data := t.data[t.HeaderSize():]
	qrowDataPtr := uint32(0)
	for i, col := range columns {
		if t.IsNull(column.Order(i)) {
			qrow.ptrs[i+1] = math.MaxUint32
			continue
		}

		valLen := col.CType().Copy(qrow.data, data)
		data = data[valLen:]
		qrowDataPtr += uint32(valLen)
		qrow.ptrs[i+1] = qrowDataPtr
	}

	for i := len(columns); i > 0; i-- {
		if qrow.ptrs[i] == math.MaxUint32 {
			qrow.ptrs[i] = qrow.ptrs[i+1]
		}
	}

	return qrow
}
