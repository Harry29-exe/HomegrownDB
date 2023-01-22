package data

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/strutils"
	"HomegrownDB/dbsystem/relation/table/column"
	page "HomegrownDB/dbsystem/storage/page/internal"
	"HomegrownDB/dbsystem/tx"
	"encoding/binary"
	"fmt"
	"strings"
)

func NewTuple(values [][]byte, pattern TuplePattern, tx tx.Tx) Tuple {
	headerLen := int(toNullBitmap + pattern.BitmapLen)
	tupleLen := headerLen
	for _, value := range values {
		tupleLen += len(value)
	}
	tuple := Tuple{
		bytes:   make([]byte, tupleLen),
		pattern: pattern,
	}

	if tx != nil {
		txId := tx.TxID()
		tuple.SetModifiedByTx(txId)
		tuple.SetCreatedByTx(txId)
	}

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

var _ WTuple = Tuple{}

type Tuple struct {
	bytes []byte

	pattern TuplePattern
}

// TID tuple id composed of Id and InPage
type TID struct {
	PageId     page.Id
	TupleIndex TupleIndex
}

type TupleIndex = uint16

const TupleIndexSize = 2

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
	subsequent := t.bytes[toNullBitmap+t.pattern.BitmapLen:]
	for i := uint16(0); i < id; i++ {
		if t.IsNull(id) {
			continue
		}
		subsequent = t.pattern.Columns[i].CType.
			Skip(subsequent)
	}

	return t.pattern.Columns[id].CType.Value(subsequent)
}

func (t Tuple) Data() []byte {
	return t.bytes[t.HeaderSize():]
}

func (t Tuple) DataSize() int {
	return len(t.bytes) - int(t.pattern.BitmapLen+toNullBitmap)
}

func (t Tuple) HeaderSize() int {
	return int(t.pattern.BitmapLen + toNullBitmap)
}

func (t Tuple) TupleSize() int {
	return len(t.bytes)
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
	return t.bytes[toPageId : toPageId+page.IdSize+TupleIndexSize]
}

func (t Tuple) tidPageIdSlice() []byte {
	return t.bytes[toPageId : toPageId+page.IdSize]
}

func (t Tuple) tidTupleIndexSlice() []byte {
	return t.bytes[toTupleIndex : toTupleIndex+TupleIndexSize]
}

func (t Tuple) NullBitmapSlice() []byte {
	if len(t.pattern.Columns) == 0 {
		panic("Tuple table is nil")
	}

	length := t.pattern.BitmapLen
	return t.bytes[toNullBitmap : toNullBitmap+length]
}

var nullBitmapMasks = [8]byte{
	1, 2, 4, 8,
	16, 32, 64, 128,
}

// -------------------------
//      Debug
// -------------------------

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

	for i := 0; i < len(tuple.pattern.Columns); i++ {
		byteIndex := uint16(i / 8)
		bitIndex := uint16(i) - byteIndex*8
		bit := bparse.Bit.GetBit(tuple.bytes[toNullBitmap+byteIndex], uint8(bitIndex))
		bitValue := uint8(0)
		if bit > 0 {
			bitValue = 1
		}

		builder.WriteString(fmt.Sprintf("| %d: %d ", i, bitValue))
	}
	builder.WriteString("|")

	arr.Add(builder.String())
}

func (t tupleDebugger) stringifyColumnValues(tuple Tuple, arr *strutils.StrArray) {
	pattern := tuple.pattern

	nextData := tuple.bytes[toNullBitmap+pattern.BitmapLen:]
	var value []byte
	for i := 0; i < len(pattern.Columns); i++ {
		col := pattern.Columns[i]
		if tuple.IsNull(column.Order(i)) {
			arr.FormatAndAdd("%s: null", col.Name)
			continue
		}

		value, nextData = col.CType.ValueAndSkip(nextData)

		arr.FormatAndAdd("%s: %s", col.Name, col.CType.ToStr(value))
	}
}
