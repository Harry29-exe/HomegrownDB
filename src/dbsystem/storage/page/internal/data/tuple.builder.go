package data

import (
	"HomegrownDB/common/datastructs/bitmap"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/tx"
	"bytes"
	"errors"
	"log"
)

type TupleBuilder interface {
	Init(pattern TuplePattern)
	WriteValue(value hgtype.Value) error
	VolatileTuple(tx tx.Tx, command uint16) Tuple
}

func NewTupleBuilder(buff bytes.Buffer) {

}

type tupleBuilder struct {
	pattern TuplePattern

	nullBitmap    bitmap.Bitmap
	dataBuff      bytes.Buffer
	valuesWritten int
}

func (t tupleBuilder) Init(pattern TuplePattern) {
	if t.pattern.BitmapLen == 0 {
		log.Panicf("tuple builder can not be initialized twice\n")
	}
	t.pattern = pattern
	t.nullBitmap = bitmap.New(int(pattern.BitmapLen))
	t.dataBuff.Write(make([]byte, toNullBitmap+pattern.BitmapLen))
}

var _ TupleBuilder = &tupleBuilder{}

func (t tupleBuilder) WriteValue(value hgtype.Value) error {
	col := t.pattern.Columns[t.valuesWritten]
	validateResult := col.Type.Validate(value)
	switch validateResult.Status {

	case hgtype.ValidateOk:
		err := col.Type.WriteValue(&t.dataBuff, value)
		if err != nil {
			return err
		}
		return nil
	case hgtype.ValidateConv:
		panic("types conv is not supported yet")
	default:
		if validateResult.Reason != nil {
			return validateResult.Reason
		}
		return errors.New("value not valid")
	}
}

func (t tupleBuilder) VolatileTuple(tx tx.Tx, commands uint16) Tuple {
	tupleData := t.dataBuff.Bytes()
	copy(tupleData[toNullBitmap:], t.nullBitmap.Bytes())

	tuple := Tuple{
		bytes:   tupleData,
		pattern: TuplePattern{},
	}
	tuple.SetCreatedByTx(tx.TxID())
	tuple.SetModifiedByTx(tx.TxID())
	tuple.SetTxCommandCounter(commands)

	return tuple
}
