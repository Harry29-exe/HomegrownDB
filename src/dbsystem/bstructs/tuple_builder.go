package bstructs

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"bytes"
)

// CreateTuple creates new TupleToSave from given columnValues and transaction context,
// Tuple inside is not initialized i.e. it does not have TID (tuple identifier) and ids of
// objects stored outside tuple should be saved to Tuple
func CreateTuple(tableDef table.Definition, columnValues map[string]any, txCtx tx.Context) (TupleToSave, error) {
	builder := tupleBuilder{}
	return builder.Create(tableDef, columnValues, txCtx)
}

type TupleToSave struct {
	Tuple           Tuple
	BgValuesToSave  []BgValueToSave
	LobValuesToSave []LobValueToSave
}

// BgValueToSave value to save in background table see documentation/storage_types.svg
type BgValueToSave struct {
	Value      []byte

}

type LobValueToSave struct {
	Value      []byte
	PtrInTuple InTuplePtr // PtrInTuple ptr to place in tuple where should be putted id to lob
}

type tupleBuilder struct {
	table        table.Definition
	sortedValues []any

	buffer    bytes.Buffer
	bufferPtr InTuplePtr

	lobs     [][]byte
	bgValues [][]byte
}

func (tb tupleBuilder) Create(tableDef table.Definition, columnValues map[string]any, txContext tx.Context) (TupleToSave, error) {
	tb.sortMapValues(tableDef, columnValues)
	tb.initTupleBuffer()
	tb.createNullBitmap()
	tb.bufferPtr = InTuplePtr(tb.buffer.Len())

	tb.serializeColumnValues()

	tuple := Tuple{
		data:  tb.buffer.Bytes(),
		table: tableDef,
	}
	tb.initTupleWithTxContext(tuple, txContext)

	return TupleToSave{
		Tuple: tuple,
		BgValuesToSave: tb.bgValues
	}tuple, nil
}

func (tb tupleBuilder) sortMapValues(tableDef table.Definition, columnValues map[string]any) {
	tb.sortedValues = make([]any, tableDef.ColumnCount())

	for i := uint16(0); i < tableDef.ColumnCount(); i++ {
		tb.sortedValues[i] = columnValues[tableDef.ColumnName(i)]
	}
}

func (tb tupleBuilder) initTupleBuffer() {
	tb.buffer = bytes.Buffer{}
	tb.buffer.Write(make([]byte, tupleHeaderSize))
}

func (tb tupleBuilder) createNullBitmap() {
	bitmapLen := tb.table.BitmapLen()
	nullBitmap := make([]byte, bitmapLen)
	colCounter := 0
	for currentByte := uint16(0); currentByte < bitmapLen; currentByte++ {
		for bit := uint8(0); bit < 8; bit++ {
			if tb.sortedValues[colCounter] != nil {
				nullBitmap[currentByte] = bparse.Bit.
					SetBit(nullBitmap[currentByte], bit)
			}

			colCounter++
		}
	}

	tb.buffer.Write(nullBitmap)
}

func (tb tupleBuilder) serializeColumnValues() {
	serializers := tb.table.AllColumnSerializer()
	for i, value := range tb.sortedValues {
		if value != nil {
			tb.serializeValue(value, serializers[i])
		} else {
			tb.table.
		}
	}
}

func (tb tupleBuilder) serializeValue(value any, serializer column.DataSerializer) {
	data, err := serializer.SerializeValue(value)
	if err != nil {
		panic(err.Error())
	}

	tb.buffer.Write(data.DataInTuple())
	if data.StorePlace() == column.StoreInBackground {
		tb.bgValues = append(tb.bgValues, data.Data())
	} else if data.StorePlace() == column.StoreInLob {
		tb.lobs = append(tb.lobs, data.Data())
	}
}

func (tb tupleBuilder) initTupleWithTxContext(tuple Tuple, ctx tx.Context) {
	tuple.SetCreatedByTx(ctx.TxId())
	tuple.SetTxCommandCounter(ctx.CommandExecuted())
}
