package page

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/io/bparse"
	"encoding/binary"
)

type Tuple struct {
	data []byte

	table table.Definition
}

// TID tuple id composed of Id and InPage
type TID struct {
	PageId     Id
	TupleIndex TupleIndex
}

type TupleIndex = uint16

const TupleIndexSize = 2

type InTuplePtr = uint16

func (t Tuple) CreatedByTx() tx.Id {
	return tx.Id(bparse.Parse.UInt4(t.data[toTxId:]))
}

func (t Tuple) ModifiedByTx() tx.Id {
	return tx.Id(bparse.Parse.UInt4(t.data[toModifiedTxId:]))
}

func (t Tuple) TxCommandCounter() uint16 {
	return bparse.Parse.UInt2(t.data[toTxId:])
}

// TID returns TID of next version of this tuple or
// TID of this tuple if its newest version
func (t Tuple) TID() TID {
	return TID{
		PageId:     bparse.Parse.UInt4(t.data[toPageId:]),
		TupleIndex: bparse.Parse.UInt2(t.data[toTupleIndex:]),
	}
}

func (t Tuple) IsNull(id column.OrderId) bool {
	var byteNumber uint16 = id / 8
	value := t.data[toNullBitmap+byteNumber]
	divRest := id % 8
	return value&nullBitmapMasks[divRest] > 0
}

func (t Tuple) ColValue(id column.OrderId) column.Value {
	subsequent := t.data[toNullBitmap+t.table.BitmapLen():]
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

func (t Tuple) SetCreatedByTx(txId tx.Id) {
	binary.LittleEndian.PutUint32(t.data[toTxId:toTxId+tx.IdSize], uint32(txId))
}

func (t Tuple) SetModifiedByTx(tx tx.Id) {
	binary.LittleEndian.PutUint32(t.modifiedByTxIdSlice(), uint32(tx))
}

func (t Tuple) SetTxCommandCounter(counter tx.Counter) {
	binary.LittleEndian.PutUint16(t.txCommandCounterSlice(), counter)
}

func (t Tuple) SetTID(tid TID) {
	binary.LittleEndian.PutUint32(t.tidPageIdSlice(), tid.PageId)
	binary.LittleEndian.PutUint16(t.tidTupleIndexSlice(), tid.TupleIndex)
}

// +++++ Binary access methods +++++

func (t Tuple) createdByTxIdSlice() []byte {
	return t.data[toTxId : toTxId+tx.IdSize]
}

func (t Tuple) modifiedByTxIdSlice() []byte {
	return t.data[toModifiedTxId : toModifiedTxId+tx.IdSize]
}

func (t Tuple) txCommandCounterSlice() []byte {
	return t.data[toTxCounter : toTxCounter+tx.CommandCounterSize]
}

func (t Tuple) tidSlice() []byte {
	return t.data[toPageId : toPageId+IdSize+TupleIndexSize]
}

func (t Tuple) tidPageIdSlice() []byte {
	return t.data[toPageId : toPageId+IdSize]
}

func (t Tuple) tidTupleIndexSlice() []byte {
	return t.data[toTupleIndex : toTupleIndex+TupleIndexSize]
}

func (t Tuple) NullBitmapSlice() []byte {
	if t.table == nil {
		panic("Tuple table is nil")
	}

	length := t.table.BitmapLen()
	return t.data[toNullBitmap : toNullBitmap+length]
}

func (t Tuple) LOBBitmap() {

}

var nullBitmapMasks = [8]byte{
	1, 2, 4, 8,
	16, 32, 64, 128,
}
