package dbbs

import (
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/storage/page"
	"math"
)

// QRow - query row is representation of data (often converted from tuple or
// concatenation of tuples) during data processing phase
type QRow struct {
	ptrs    []uint32
	data    []byte
	Pattern []ctype.CType
}

func NewQRowFromTuple(tuple page.RTuple) QRow {
	t := tuple.(page.Tuple)
	table := t.Table()
	columns := table.Columns()

	headerLen := 4 * int(table.ColumnCount()+1)
	qrow := QRow{
		data:    make([]byte, t.DataSize()),
		ptrs:    make([]uint32, headerLen),
		Pattern: table.CTypePattern(),
	}

	data := t.Data()
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

func (qr QRow) Value(index uint32) []byte {
	//todo implement me
	panic("Not implemented")
}
