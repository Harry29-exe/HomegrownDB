package buffer_test

import (
	"HomegrownDB/dbsystem/buffer"
	"HomegrownDB/dbsystem/schema/column/factory"
	"HomegrownDB/dbsystem/schema/column/types"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/tests"
	"fmt"
	"testing"
)

func TestCreateEmptyPage(t *testing.T) {
	tableDef := table.NewDefinition(
		"test_table", 0, 0)

	newPage := buffer.CreateEmptyPage(tableDef)

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
	table := pUtils.testTable()
	page := buffer.CreateEmptyPage(table)

	txCtx1 := tx.NewContext(1)
	tupleToSave, err := buffer.TupleBuilder.Create(table, pUtils.colValues1(), txCtx1)
	if err != nil {
		t.Error(err.Error())
	}
	tuple := tupleToSave.Tuple

	err = page.InsertTuple(tuple)
	if err != nil {
		t.Errorf("%e", err)
	}

	tests.AssertEq(page.TupleCount(), 1, t)
	tests.AssertEqArray(tuple.Data(), page.Tuple(0).Data(), t)
}

var pUtils = pageUtils{}

type pageUtils struct{}

func (u pageUtils) testTable() table.Definition {
	def := table.NewDefinition("test_table", 12, 20)
	def.AddColumn(factory.CreateDefinition(
		factory.ArgsBuilder().
			Name("col1").
			Type(types.Int2).
			Build(),
	))
	def.AddColumn(factory.CreateDefinition(
		factory.ArgsBuilder().
			Name("col2").
			Type(types.Int2).
			Build(),
	))

	return def
}

func (u pageUtils) colValues1() map[string]any {
	return map[string]any{
		"col1": 1,
		"col2": 2,
	}
}
