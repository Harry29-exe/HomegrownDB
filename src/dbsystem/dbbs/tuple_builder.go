package dbbs

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
)

func NewTuple(values [][]byte, table table.Definition, tx tx.Ctx) Tuple {
	headerLen := int(toNullBitmap + table.BitmapLen())
	tupleLen := headerLen
	for _, value := range values {
		tupleLen += len(value)
	}
	tuple := Tuple{
		data:  make([]byte, tupleLen),
		table: table,
	}

	txId := tx.Info.TxId()
	tuple.SetModifiedByTx(txId)
	tuple.SetCreatedByTx(txId)

	tupleData := tuple.data[headerLen:]
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
