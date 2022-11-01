package qrow_test

import (
	"HomegrownDB/backend/qrow"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/ttable1"
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/tpage"
	"HomegrownDB/dbsystem/tx"
	"testing"
)

func TestNewRow(t *testing.T) {
	tableDef := ttable1.Def(t)
	testTuple, err := tpage.NewTestTuple(tableDef, map[string][]byte{
		ttable1.C0AwesomeKey:  convInput(int64(24), tableDef.Column(ttable1.C0AwesomeKeyOrder).Type()),
		ttable1.C1NullableCol: nil,
		ttable1.C2NonNullColl: convInput(int64(43), tableDef.Column(ttable1.C2NonNullCollOrder).Type()),
	}, tx.NewInfoCtx(29))
	if err != nil {
		panic(err.Error())
	}

	buffer := qrow.NewSlotBuffer(10_000)
	holder := qrow.NewBaseRowHolder(buffer, []table.Definition{tableDef})
	dataRow := qrow.NewRow([]tpage.Tuple{testTuple.Tuple}, holder)

	assert.EqArray(
		tpage.TupleHelper.GetValueByName(testTuple.Tuple, tableDef, ttable1.C0AwesomeKey),
		dataRow.GetField(ttable1.C0AwesomeKeyOrder),
		t,
	)
	assert.EqArray(
		tpage.TupleHelper.GetValueByName(testTuple.Tuple, tableDef, ttable1.C1NullableCol),
		dataRow.GetField(ttable1.C1NullableColOrder),
		t,
	)
	assert.EqArray(
		tpage.TupleHelper.GetValueByName(testTuple.Tuple, tableDef, ttable1.C2NonNullColl),
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
