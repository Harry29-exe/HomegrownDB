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

type TupleId struct {
	PageId        uint32
	InPagePointer InPagePointer
}

type TupleIndex = uint16

type InTuplePtr = uint16

func (t Tuple) CreatedByTx() tx.Id {
	return bparse.Parse.UInt4(t.data[tupleTxId:])
}

func (t Tuple) ModifiedByTx() tx.Id {
	return bparse.Parse.UInt4(t.data[tupleUpdateTxId:])
}

func (t Tuple) TxCommandCounter() uint16 {
	return bparse.Parse.UInt2(t.data[tupleTxId:])
}

// TID returns TupleId of next version of this tuple or
// TupleId of this tuple if its newest version
func (t Tuple) TID() TupleId {
	return TupleId{
		PageId:        bparse.Parse.UInt4(t.data[tuplePageId:]),
		InPagePointer: bparse.Parse.UInt2(t.data[tupleLinePointer:]),
	}
}

func (t Tuple) IsNull(id column.OrderId) bool {
	var byteNumber uint16 = id / 8
	value := t.data[tupleNullBitmap+byteNumber]
	divRest := id % 8
	return value&nullBitmapMasks[divRest] > 0
}

func (t Tuple) ColValue(id column.OrderId) column.Value {
	subsequent := t.data[tupleNullBitmap+t.table.NullBitmapLen():]
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
	tupleTxId        = 0                    // offset to created by tx_id field
	tupleUpdateTxId  = 4 + tupleTxId        // offset to created/modified by tx_id field
	tupleTxCounter   = 4 + tupleUpdateTxId  // offset to amount of command executed by TxId
	tuplePageId      = 2 + tupleTxCounter   //offset to pageId where next version of this tuple can be found
	tupleLinePointer = 4 + tuplePageId      // offset to line number of next version of this tuple
	tupleNullBitmap  = tupleLinePointer + 2 // offset to start of null bitmap
)
