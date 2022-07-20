package bdata_test

import (
	"HomegrownDB/common/random"
	"HomegrownDB/common/tests"
	"HomegrownDB/common/tests/tutils"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/ctypes"
	"HomegrownDB/dbsystem/schema/column/factory"
	"HomegrownDB/dbsystem/schema/table"
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

	tests.ShouldPanic(
		func() {
			newPage.Tuple(0)
		},
		"Newly created tuple returned non existing tuple with index 0", t)

	tests.ShouldPanic(
		func() {
			newPage.UpdateTuple(0, make([]byte, 128))
		},
		"Newly created tuple updated non existing tuple with index 0", t)
}

func TestPage_Tuple(t *testing.T) {
	table := tutils.TestTables.Table1Def()
	page := bdata.CreateEmptyPage(table)

	//txCtx1 := tx.NewContext(1)
	rand := random.NewRandom(13)
	for tupleIndex := uint16(0); tupleIndex < 20; tupleIndex++ {
		tuple := table.RandTuple(rand).Tuple
		err := page.InsertTuple(tuple.Data())
		if err != nil {
			t.Errorf("could not insert tuple nr %d because of error: %e", tupleIndex, err)
			tests.AssertEq(page.TupleCount(), tupleIndex, t)
			tests.AssertEqArray(tuple.Data(), page.Tuple(tupleIndex).Data(), t)
		}

	}
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
