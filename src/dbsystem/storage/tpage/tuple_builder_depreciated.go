package tpage

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/tx"
	"bytes"
	"fmt"
)

// NewTestTuple creates new TupleToSave from given columnValues and transaction context,
// Tuple inside is not initialized i.e. it does not have TID (tuple identifier) and ids of
// objects stored outside tuple should be saved to Tuple
func NewTestTuple(tableDef table.RDefinition, columnValues map[string][]byte, txInfo *tx.InfoCtx) (TupleToSave, error) {
	builder := tupleBuilder{table: tableDef}
	tuple, err := builder.Create(columnValues, txInfo)
	if err != nil {
		return TupleToSave{}, err
	}

	return TupleToSave{
		Tuple: tuple,
	}, nil
}

type tupleBuilder struct {
	table        table.RDefinition
	sortedValues [][]byte

	buffer bytes.Buffer
}

func (tb *tupleBuilder) Create(columnValues map[string][]byte, txInfo *tx.InfoCtx) (Tuple, error) {
	tb.sortMapValues(columnValues)
	tb.initTupleBuffer()
	tb.createNullBitmap()

	err := tb.serializeColumnValues()
	if err != nil {
		return Tuple{}, err
	}

	tuple := Tuple{
		bytes: tb.buffer.Bytes(),
		table: tb.table,
	}
	tb.initTupleWithTxContext(tuple, txInfo)

	return tuple, nil
}

func (tb *tupleBuilder) sortMapValues(columnValues map[string][]byte) {
	tb.sortedValues = make([][]byte, tb.table.ColumnCount())

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
		colDef := tableDef.Column(uint16(i))

		if value != nil {
			tb.buffer.Write(value)
		} else if !colDef.Nullable() {
			return fmt.Errorf("column %s is not nullable, so it can not accept null value",
				colDef.Name())

		}
	}

	return nil
}

func (tb *tupleBuilder) initTupleWithTxContext(tuple Tuple, txInfo *tx.InfoCtx) {
	tuple.SetCreatedByTx(txInfo.TxId())
	tuple.SetTxCommandCounter(txInfo.CommandExecuted())
}
