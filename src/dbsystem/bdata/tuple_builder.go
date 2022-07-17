package bdata

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/io/lob"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"bytes"
	"fmt"
)

// CreateTuple creates new TupleToSave from given columnValues and transaction context,
// Tuple inside is not initialized i.e. it does not have TID (tuple identifier) and ids of
// objects stored outside tuple should be saved to Tuple
func CreateTuple(tableDef table.Definition, columnValues map[string]any, txCtx tx.Context) (TupleToSave, error) {
	builder := tupleBuilder{table: tableDef}
	tuple, err := builder.Create(columnValues, txCtx)
	if err != nil {
		return TupleToSave{}, err
	}

	return TupleToSave{
		Tuple:           tuple,
		BgValuesToSave:  builder.bgValues,
		LobValuesToSave: builder.lobs,
	}, nil
}

type tupleBuilder struct {
	table        table.Definition
	sortedValues []any

	buffer bytes.Buffer

	lobs     []LobValueToSave
	bgValues []BgValueToSave
}

func (tb *tupleBuilder) Create(columnValues map[string]any, txContext tx.Context) (Tuple, error) {
	tb.sortMapValues(columnValues)
	tb.initTupleBuffer()
	tb.createNullBitmap()

	err := tb.serializeColumnValues()
	if err != nil {
		return Tuple{}, err
	}

	tuple := Tuple{
		data:  tb.buffer.Bytes(),
		table: tb.table,
	}
	tb.initTupleWithTxContext(tuple, txContext)

	return tuple, nil
}

func (tb *tupleBuilder) sortMapValues(columnValues map[string]any) {
	tb.sortedValues = make([]any, tb.table.ColumnCount())

	for i := uint16(0); i < tb.table.ColumnCount(); i++ {
		tb.sortedValues[i] = columnValues[tb.table.ColumnName(i)]
	}
}

func (tb *tupleBuilder) initTupleBuffer() {
	tb.buffer = bytes.Buffer{}
	tb.buffer.Write(make([]byte, tupleHeaderSize))
}

func (tb *tupleBuilder) createNullBitmap() {
	bitmapLen := tb.table.BitmapLen()
	nullBitmap := make([]byte, bitmapLen)
	colCounter := 0
	currentByte := uint16(0)
	for ; currentByte < bitmapLen-1; currentByte++ {
		for bit := uint8(0); bit < 8; bit++ {
			if tb.sortedValues[colCounter] != nil {
				nullBitmap[currentByte] = bparse.Bit.
					SetBit(nullBitmap[currentByte], bit)
			}

			colCounter++
		}
	}

	colCounter = 0
	bitsInLastByte := uint8(tb.table.ColumnCount() - (bitmapLen-1)*8)
	for bit := uint8(0); bit < bitsInLastByte; bit++ {
		if tb.sortedValues[colCounter] != nil {
			nullBitmap[currentByte] = bparse.Bit.
				SetBit(nullBitmap[currentByte], bit)
		}

		colCounter++
	}

	tb.buffer.Write(nullBitmap)
}

func (tb *tupleBuilder) serializeColumnValues() error {
	tableDef := tb.table
	for i, value := range tb.sortedValues {
		colDef := tableDef.GetColumn(uint16(i))

		if value != nil {
			serializer := colDef.DataSerializer()
			data, err := serializer.SerializeValue(value)
			if err != nil {
				return err
			}

			tb.saveData(data, colDef)
		} else if !colDef.Nullable() {
			return fmt.Errorf("column %s is not nullable, so it can not accept null value",
				colDef.Name())

		}
	}

	return nil
}

func (tb *tupleBuilder) saveData(data column.DataToSave, col column.Definition) {
	tb.buffer.Write(data.DataInTuple())
	if data.StorePlace() == column.StoreInBackground {
		tb.bgValues = append(tb.bgValues,
			BgValueToSave{data.Data(), col.GetColumnId()})

	} else if data.StorePlace() == column.StoreInLob {
		lobId := lob.IdCounter.NextId()
		tb.lobs = append(tb.lobs, LobValueToSave{data.Data(), lobId})
	}
}

func (tb *tupleBuilder) initTupleWithTxContext(tuple Tuple, ctx tx.Context) {
	tuple.SetCreatedByTx(ctx.TxId())
	tuple.SetTxCommandCounter(ctx.CommandExecuted())
}
