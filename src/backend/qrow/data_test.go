package qrow_test

import (
	"HomegrownDB/backend/qrow"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/testtable/ttable1"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"testing"
)

func TestNewRow(t *testing.T) {
	tableDef := ttable1.Def()
	testTuple, err := bdata.CreateTuple(tableDef, map[string][]byte{
		ttable1.C0AwesomeKey:  convInput(int64(24), tableDef.Column(ttable1.C0AwesomeKeyOrder).Type()),
		ttable1.C1NullableCol: nil,
		ttable1.C2NonNullColl: convInput(int64(43), tableDef.Column(ttable1.C2NonNullCollOrder).Type()),
	}, tx.NewInfoCtx(29))
	if err != nil {
		panic(err.Error())
	}

	buffer := qrow.NewSlotBuffer(10_000)
	holder := qrow.NewBaseRowHolder(buffer, []table.Definition{tableDef})
	dataRow := qrow.NewRow([]bdata.Tuple{testTuple.Tuple}, holder)

	assert.EqArray(
		bdata.TupleHelper.GetValueByName(testTuple.Tuple, tableDef, ttable1.C0AwesomeKey),
		dataRow.GetField(ttable1.C0AwesomeKeyOrder),
		t,
	)
	assert.EqArray(
		bdata.TupleHelper.GetValueByName(testTuple.Tuple, tableDef, ttable1.C1NullableCol),
		dataRow.GetField(ttable1.C1NullableColOrder),
		t,
	)
	assert.EqArray(
		bdata.TupleHelper.GetValueByName(testTuple.Tuple, tableDef, ttable1.C2NonNullColl),
		dataRow.GetField(ttable1.C2NonNullCollOrder),
		t,
	)
}

func convInput(input any, cType ctype.Type) []byte {
	v, err := ctype.ConvInput(input, cType)
	if err != nil {
		panic("unexpected error, most likely test is invalid")
	}
	return v
}
