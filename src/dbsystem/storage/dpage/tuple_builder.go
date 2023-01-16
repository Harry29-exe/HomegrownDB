package dpage

import (
	"HomegrownDB/dbsystem/relation/table/column"
	"HomegrownDB/dbsystem/tx"
)

func NewTuple(values [][]byte, pattern *TuplePattern, tx tx.Tx) Tuple {
	headerLen := int(toNullBitmap + pattern.BitmapLen)
	tupleLen := headerLen
	for _, value := range values {
		tupleLen += len(value)
	}
	tuple := Tuple{
		bytes:   make([]byte, tupleLen),
		pattern: pattern,
	}

	txId := tx.TxID()
	tuple.SetModifiedByTx(txId)
	tuple.SetCreatedByTx(txId)

	tupleData := tuple.bytes[headerLen:]
	var copiedBytes int
	for i, value := range values {
		if value == nil {
			tuple.SetIsNull(column.Order(i))
			continue
		}

		copiedBytes = copy(tupleData, value)
		tupleData = tupleData[copiedBytes:]
	}

	return tuple
}
