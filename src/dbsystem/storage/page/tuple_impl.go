package page

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/strutils"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"encoding/binary"
	"fmt"
	"strings"
)

var _ WTuple = Tuple{}

type Tuple struct {
	bytes []byte

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

func (t Tuple) Bytes() []byte {
	return t.bytes
}

func (t Tuple) CreatedByTx() tx.Id {
	return tx.Id(bparse.Parse.UInt4(t.bytes[toTxId:]))
}

func (t Tuple) ModifiedByTx() tx.Id {
	return tx.Id(bparse.Parse.UInt4(t.bytes[toModifiedTxId:]))
}

func (t Tuple) TxCommandCounter() uint16 {
	return bparse.Parse.UInt2(t.bytes[toTxId:])
}

// TID returns TID of next version of this tuple or
// TID of this tuple if its newest version
func (t Tuple) TID() TID {
	return TID{
		PageId:     bparse.Parse.UInt4(t.bytes[toPageId:]),
		TupleIndex: bparse.Parse.UInt2(t.bytes[toTupleIndex:]),
	}
}

func (t Tuple) IsNull(id column.Order) bool {
	var byteNumber = id / 8
	value := t.bytes[toNullBitmap+byteNumber]
	divRest := id % 8
	return value&nullBitmapMasks[divRest] == 0
}

func (t Tuple) SetIsNull(id column.Order) {
	byteNumber := id / 8
	bytePos := toNullBitmap + byteNumber

	divRest := id % 8
	t.bytes[bytePos] = bparse.Bit.SetBit(t.bytes[bytePos], uint8(divRest))
}

func (t Tuple) SetIsNotNull(id column.Order) {
	byteNumber := id / 8
	bytePos := toNullBitmap + byteNumber

	divRest := id % 8
	t.bytes[bytePos] = bparse.Bit.ClearBit(t.bytes[bytePos], uint8(divRest))
}

func (t Tuple) ColValue(id column.Order) []byte {
	subsequent := t.bytes[toNullBitmap+t.table.BitmapLen():]
	for i := uint16(0); i < id; i++ {
		if t.IsNull(id) {
			continue
		}
		subsequent = t.table.
			ColumnType(i).
			Skip(subsequent)
	}

	return t.table.ColumnType(id).Value(subsequent)
}

func (t Tuple) Data() []byte {
	return t.bytes[t.HeaderSize():]
}

func (t Tuple) DataSize() int {
	return len(t.bytes) - int(t.table.BitmapLen()+toNullBitmap)
}

func (t Tuple) Table() table.Definition {
	return t.table
}

func (t Tuple) HeaderSize() int {
	return int(t.table.BitmapLen() + toNullBitmap)
}

func (t Tuple) SetCreatedByTx(txId tx.Id) {
	binary.BigEndian.PutUint32(t.bytes[toTxId:toTxId+tx.IdSize], uint32(txId))
}

func (t Tuple) SetModifiedByTx(tx tx.Id) {
	binary.BigEndian.PutUint32(t.modifiedByTxIdSlice(), uint32(tx))
}

func (t Tuple) SetTxCommandCounter(counter tx.Counter) {
	binary.BigEndian.PutUint16(t.txCommandCounterSlice(), counter)
}

func (t Tuple) SetTID(tid TID) {
	binary.BigEndian.PutUint32(t.tidPageIdSlice(), tid.PageId)
	binary.BigEndian.PutUint16(t.tidTupleIndexSlice(), tid.TupleIndex)
}

// +++++ Binary access methods +++++

func (t Tuple) createdByTxIdSlice() []byte {
	return t.bytes[toTxId : toTxId+tx.IdSize]
}

func (t Tuple) modifiedByTxIdSlice() []byte {
	return t.bytes[toModifiedTxId : toModifiedTxId+tx.IdSize]
}

func (t Tuple) txCommandCounterSlice() []byte {
	return t.bytes[toTxCounter : toTxCounter+tx.CommandCounterSize]
}

func (t Tuple) tidSlice() []byte {
	return t.bytes[toPageId : toPageId+IdSize+TupleIndexSize]
}

func (t Tuple) tidPageIdSlice() []byte {
	return t.bytes[toPageId : toPageId+IdSize]
}

func (t Tuple) tidTupleIndexSlice() []byte {
	return t.bytes[toTupleIndex : toTupleIndex+TupleIndexSize]
}

func (t Tuple) NullBitmapSlice() []byte {
	if t.table == nil {
		panic("Tuple table is nil")
	}

	length := t.table.BitmapLen()
	return t.bytes[toNullBitmap : toNullBitmap+length]
}

var nullBitmapMasks = [8]byte{
	1, 2, 4, 8,
	16, 32, 64, 128,
}

// +++++ Debug +++++

// TupleDebugger helps with debugging of Tuple structure
var TupleDebugger = tupleDebugger{}

type tupleDebugger struct{}

func (t tupleDebugger) TupleDescription(tuple Tuple) []string {
	strArr := &strutils.StrArray{}

	strArr.FormatAndAdd("Created by %d", tuple.CreatedByTx())
	strArr.FormatAndAdd("Modified by %d", tuple.ModifiedByTx())
	strArr.FormatAndAdd("Command executed before %d", tuple.TxCommandCounter())
	tid := tuple.TID()
	strArr.FormatAndAdd("TID: (PageId: %d, TupleIndex: %d)", tid.PageId, tid.TupleIndex)

	t.stringifyNullBitmap(tuple, strArr)
	t.stringifyColumnValues(tuple, strArr)

	//todo col values
	return strArr.Array
}

func (t tupleDebugger) stringifyNullBitmap(tuple Tuple, arr *strutils.StrArray) {
	builder := strings.Builder{}
	builder.WriteString("NullBitmap: ")
	for _, bitmapByte := range tuple.NullBitmapSlice() {
		builder.WriteString(fmt.Sprintf("%08b", bitmapByte))
	}
	builder.WriteRune('\n')

	for i := column.Order(0); i < tuple.table.ColumnCount(); i++ {
		col := tuple.table.Column(i)
		if !col.Nullable() {
			builder.WriteString(fmt.Sprintf("| %s: %d ", col.Name(), -1))
			continue
		}

		byteIndex := i / 8
		bitIndex := i - byteIndex*8
		bit := bparse.Bit.GetBit(tuple.bytes[toNullBitmap+byteIndex], uint8(bitIndex))
		bitValue := uint8(0)
		if bit > 0 {
			bitValue = 1
		}

		builder.WriteString(fmt.Sprintf("| %s: %d ", col.Name(), bitValue))
	}
	builder.WriteString("|")

	arr.Add(builder.String())
}

func (t tupleDebugger) stringifyColumnValues(tuple Tuple, arr *strutils.StrArray) {
	tabDef := tuple.table

	nextData := tuple.bytes[toNullBitmap+tabDef.BitmapLen():]
	var value []byte
	for i := column.Order(0); i < tabDef.ColumnCount(); i++ {
		col := tabDef.Column(i)
		if tuple.IsNull(i) {
			arr.FormatAndAdd("%s: null", col.Name())
			continue
		}

		parser := col.CType()

		value, nextData = parser.ValueAndSkip(nextData)

		arr.FormatAndAdd("%s: %s", col.Name(), parser.ToStr(value))
	}
}
