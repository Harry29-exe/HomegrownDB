package bdata_test

import (
	tests2 "HomegrownDB/common/tests"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/ctypes"
	"HomegrownDB/dbsystem/schema/column/factory"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"fmt"
	"testing"
)

func TestCreateEmptyPage(t *testing.T) {
	tableDef := table.NewDefinition(
		"test_table")
	tableDef.SetTableId(32)
	tableDef.SetObjectId(12)

	newPage := bdata.CreateEmptyPage(tableDef)

	if newPage.TupleCount() != 0 {
		errMsg := fmt.Sprintf("new page has tuple count different than 0,\n page: %#v", newPage)
		t.Error(errMsg)
	}

	tests2.ShouldPanic(
		func() {
			newPage.Tuple(0)
		},
		"Newly created tuple returned non existing tuple with index 0", t)

	tests2.ShouldPanic(
		func() {
			newPage.UpdateTuple(0, make([]byte, 128))
		},
		"Newly created tuple updated non existing tuple with index 0", t)
}

func TestPage_Tuple(t *testing.T) {
	table := pUtils.testTable()
	page := bdata.CreateEmptyPage(table)

	txCtx1 := tx.NewContext(1)
	tupleToSave, err := bdata.CreateTuple(table, pUtils.colValues1(), txCtx1)
	if err != nil {
		t.Error(err.Error())
	}
	tuple := tupleToSave.Tuple

	err = page.InsertTuple(tuple.Data())
	if err != nil {
		t.Errorf("%e", err)
	}

	tests2.AssertEq(page.TupleCount(), 1, t)
	tests2.AssertEqArray(tuple.Data(), page.Tuple(0).Data(), t)
}

var pUtils = pageUtils{}

type pageUtils struct{}

func (u pageUtils) testTable() table.Definition {
	tableDef := table.NewDefinition("test_table")
	tableDef.SetTableId(12)
	tableDef.SetObjectId(20)

	tableDef.AddColumn(factory.CreateDefinition(
		column.ArgsBuilder("col1", ctypes.Int2).Build(),
	))
	tableDef.AddColumn(factory.CreateDefinition(
		column.ArgsBuilder("col2", ctypes.Int2).Build(),
	))

	return tableDef
}

func (u pageUtils) colValues1() map[string]any {
	return map[string]any{
		"col1": 1,
		"col2": 2,
	}
}
