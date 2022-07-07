package bdata

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/tx"
)

type RTuple interface {
	CreatedByTx() tx.Id
	ModifiedByTx() tx.Id
	TxCommandCounter() uint16
	TID() TID
	IsNull(id column.OrderId) bool
	ColValue(id column.OrderId) column.Value
}

type WTuple interface {
	RTuple
	SetCreatedByTx(txId tx.Id)
	SetModifiedByTx(tx tx.Id)
	SetTxCommandCounter(counter uint16)
	SetTID(tid TID)
}

type TupleOffset = uint16

const (
	sizeOfTxId         = tx.IdSize
	sizeOfModifiedTxId = tx.IdSize
	sizeOfTxCounter    = tx.CommandCounterSize
	sizeOfPageId       = PageIdSize
	sizeOfTupleIndex   = TupleIndexSize
)

const (
	toTxId         TupleOffset = 0                                   // offset to created by tx_id field
	toModifiedTxId             = sizeOfTxId + toTxId                 // offset to created/modified by tx_id field
	toTxCounter                = sizeOfModifiedTxId + toModifiedTxId // offset to amount of command executed by TxId
	toPageId                   = sizeOfTxCounter + toTxCounter       // offset to pageId where next version of this tuple can be found
	toTupleIndex               = sizeOfPageId + toPageId             // offset to line number of next version of this tuple
	toNullBitmap               = sizeOfTupleIndex + toTupleIndex     // offset to start of null bitmap
)

const (
	tupleHeaderSize = toNullBitmap // size in bytes of tuple header data
)
