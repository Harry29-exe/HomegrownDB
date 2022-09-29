package qrow_test

import (
	"HomegrownDB/backend/qrow"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"testing"
)

func TestNewRow(t *testing.T) {
	tableDef := tutils.TestTables.Table1Def()
	testTuple, err := bdata.CreateTuple(tableDef, map[string][]byte{
		tutils.Table1.AwesomeKey:  convInput(24, tableDef.Column(tutils.Table1.AwesomeKeyId).Type()),
		tutils.Table1.NullableCol: nil,
		tutils.Table1.NonNullColl: convInput(43, tableDef.Column(tutils.Table1.NonNullCollId).Type()),
	}, tx.NewInfoCtx(29))
	if err != nil {
		panic(err.Error())
	}

	buffer := qrow.NewSlotBuffer(10_000)
	holder := qrow.NewBaseRowHolder(buffer, []table.Definition{tableDef})
	dataRow := qrow.NewRow([]bdata.Tuple{testTuple.Tuple}, holder)

	assert.EqArray(
		bdata.TupleHelper.GetValueByName(testTuple.Tuple, tableDef, tutils.Table1.AwesomeKey),
		dataRow.GetField(tutils.Table1.AwesomeKeyId),
		t,
	)
	assert.EqArray(
		bdata.TupleHelper.GetValueByName(testTuple.Tuple, tableDef, tutils.Table1.NullableCol),
		dataRow.GetField(tutils.Table1.NullableColId),
		t,
	)
	assert.EqArray(
		bdata.TupleHelper.GetValueByName(testTuple.Tuple, tableDef, tutils.Table1.NonNullColl),
		dataRow.GetField(tutils.Table1.NonNullCollId),
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
