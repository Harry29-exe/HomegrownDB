package query

import (
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/storage/tpage"
	"math"
)

// QRow - query row is representation of data (often converted from tuple or
// concatenation of tuples) during data processing phase
type QRow struct {
	ptrs    []uint32
	data    []byte
	Pattern []ctype.CType
}

func NewQRowFromTuple(tuple tpage.RTuple) QRow {
	t := tuple.(tpage.Tuple)
	table := t.Table()
	columns := table.Columns()

	headerLen := table.ColumnCount() + 1
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

func NewQRowFromValues(values [][]byte, valuesCTypes []ctype.CType) QRow {
	dataLen := uint32(0)
	qRowPtrs := make([]uint32, len(values))
	for i, value := range values {
		qRowPtrs[i] = dataLen
		dataLen += uint32(len(value))
	}
	qRowPtrs[len(values)] = dataLen

	qRowData := make([]byte, dataLen)
	for i := 0; i < len(values); i++ {
		copy(qRowData[qRowPtrs[i]:], values[i])
	}

	return QRow{
		ptrs:    qRowPtrs,
		data:    qRowData,
		Pattern: valuesCTypes,
	}
}

func (qr QRow) Value(index uint32) []byte {
	//todo implement me
	panic("Not implemented")
}
