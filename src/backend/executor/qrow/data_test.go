package qrow_test

import (
	qrow2 "HomegrownDB/backend/executor/qrow"
	"HomegrownDB/common/tests"
	"HomegrownDB/common/tests/tutils"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"testing"
)

func TestNewRow(t *testing.T) {
	tableDef := tutils.TestTables.Table1Def()
	testTuple, err := bdata.CreateTuple(tableDef, map[string]any{
		tutils.Table1.AwesomeKey:  24,
		tutils.Table1.NullableCol: nil,
		tutils.Table1.NonNullColl: 43,
	}, tx.NewContext(29))
	if err != nil {
		panic(err.Error())
	}

	buffer := qrow2.NewSlotBuffer(10_000)
	holder := qrow2.NewBaseRowHolder(buffer, []table.Definition{tableDef})
	dataRow := qrow2.NewRow([]bdata.Tuple{testTuple.Tuple}, holder)

	tests.AssertEqArray(
		bdata.TupleHelper.GetValueByName(testTuple.Tuple, tableDef, tutils.Table1.AwesomeKey),
		dataRow.GetField(tutils.Table1.AwesomeKeyId),
		t,
	)
	tests.AssertEqArray(
		bdata.TupleHelper.GetValueByName(testTuple.Tuple, tableDef, tutils.Table1.NullableCol),
		dataRow.GetField(tutils.Table1.NullableColId),
		t,
	)
	tests.AssertEqArray(
		bdata.TupleHelper.GetValueByName(testTuple.Tuple, tableDef, tutils.Table1.NonNullColl),
		dataRow.GetField(tutils.Table1.NonNullCollId),
		t,
	)
}
