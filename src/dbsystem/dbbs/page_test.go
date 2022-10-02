package dbbs_test

import (
	"HomegrownDB/common/random"
	"HomegrownDB/common/tests"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/testtable/ttable1"
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/schema/table"
	"fmt"
	"testing"
)

func TestCreateEmptyPage(t *testing.T) {
	tableDef := table.NewDefinition(
		"test_table")
	tableDef.SetTableId(32)
	tableDef.SetObjectId(12)

	newPage := dbbs.CreateEmptyPage(tableDef)

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
	table1 := ttable1.Def()
	page := dbbs.CreateEmptyPage(table1)

	//txCtx1 := tx.NewContext(1)
	rand := random.NewRandom(13)
	for tupleIndex := uint16(0); tupleIndex < 20; tupleIndex++ {
		tuple := table1.RandTuple(rand).Tuple
		err := page.InsertTuple(tuple.Bytes())
		if err != nil {
			t.Errorf("could not insert tuple nr %d because of error: %e", tupleIndex, err)
		}
		assert.Eq(page.TupleCount(), tupleIndex+1, t)
		assert.EqArray(tuple.Bytes(), page.Tuple(tupleIndex).Bytes(), t)
	}
}

func TestPage_DeleteTuple_FromMiddle(t *testing.T) {
	table1 := ttable1.Def()
	page := dbbs.CreateEmptyPage(table1)

	tuples := table1.PutRandomTupleToPage(10, page, random.NewRandom(23))
	page.DeleteTuple(2)
	page.DeleteTuple(8)

	assertTuplesList := []dbbs.TupleIndex{0, 1, 3, 4, 5, 6, 7, 9}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(page.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
}

func TestPage_DeleteTuple_First(t *testing.T) {
	table1 := ttable1.Def()
	page := dbbs.CreateEmptyPage(table1)

	tuples := table1.PutRandomTupleToPage(10, page, random.NewRandom(23))

	page.DeleteTuple(0)
	assertTuplesList := []dbbs.TupleIndex{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(page.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
	page.DeleteTuple(1)
	assertTuplesList = []dbbs.TupleIndex{2, 3, 4, 5, 6, 7, 8, 9}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(page.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
}

func TestPage_DeleteTuple_Last(t *testing.T) {
	table1 := ttable1.Def()
	page := dbbs.CreateEmptyPage(table1)

	tuples := table1.PutRandomTupleToPage(10, page, random.NewRandom(23))

	page.DeleteTuple(9)
	assertTuplesList := []dbbs.TupleIndex{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(page.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
	page.DeleteTuple(8)
	assertTuplesList = []dbbs.TupleIndex{0, 1, 2, 3, 4, 5, 6, 7}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(page.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
}
