package page

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/io/bparse"
)

type Tuple struct {
	data []byte
	//CreatedByTx      uint32
	//ModifiedByTx     uint32
	//TxCommandCounter uint32
	//Id               Id
	//
	//nullBitmap uint16
	//
	//columns []Column
	table table.Definition
}

type LinePointer = uint16

type TupleId struct {
	PageId      uint32
	LinePointer LinePointer
}

func (t Tuple) CreatedByTx() tx.Id {
	return bparse.Parse.UInt4(t.data[offsetTxId:])
}

func (t Tuple) ModifiedByTx() tx.Id {
	return bparse.Parse.UInt4(t.data[offsetUpdateTxId:])
}

func (t Tuple) TxCommandCounter() uint16 {
	return bparse.Parse.UInt2(t.data[offsetTxId:])
}

// TID returns TupleId of next version of this tuple or
// TupleId of this tuple if its newest version
func (t Tuple) TID() TupleId {
	return TupleId{
		PageId:      bparse.Parse.UInt4(t.data[offsetPageId:]),
		LinePointer: bparse.Parse.UInt2(t.data[offsetLinePointer:]),
	}
}

func (t Tuple) IsNull(id column.OrderId) bool {
	var byteNumber uint16 = id / 8
	value := t.data[offsetNullBitmap+byteNumber]
	divRest := id % 8
	return value&nullBitmapMasks[divRest] > 0
}

func (t Tuple) ColValue(id column.OrderId) column.Value {
	subsequent := t.data[offsetNullBitmap+t.table.NullBitmapLen():]
	for i := uint16(0); i < id; i++ {
		if t.IsNull(id) {
			continue
		}
		subsequent = t.table.
			ColumnParser(i).
			Skip(subsequent)
	}

	val, er := t.table.ColumnParser(id).Parse(subsequent)
	if er != nil {
		panic("could not parse table")
	}

	return val
}

var nullBitmapMasks = [8]byte{
	1, 2, 4, 8,
	16, 32, 64, 128,
}

const (
	offsetTxId        = 0                     // offset to created by tx_id field
	offsetUpdateTxId  = 4 + offsetTxId        // offset to created/modified by tx_id field
	offsetTxCounter   = 4 + offsetUpdateTxId  // offset to amount of command executed by TxId
	offsetPageId      = 2 + offsetTxCounter   //offset to pageId where next version of this tuple can be found
	offsetLinePointer = 4 + offsetPageId      // offset to line number of next version of this tuple
	offsetNullBitmap  = offsetLinePointer + 2 // offset to start of null bitmap
)
