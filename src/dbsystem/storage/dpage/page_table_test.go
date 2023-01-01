package dpage_test

import (
	"HomegrownDB/common/random"
	"HomegrownDB/common/tests"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/ttable1"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/tpage"
	"fmt"
	"testing"
)

func TestCreateEmptyPage(t *testing.T) {
	tableDef := table.NewDefinition(
		"test_table")
	tableDef.SetRelationID(32)
	tableDef.SetRelationID(12)

	newPage := tpage.EmptyTablePage(tableDef, t)

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
	table1 := ttable1.Def(t)
	page := tpage.EmptyTablePage(table1, t)

	//txCtx1 := tx.NewContext(1)
	table1.TUtils.SetRand(random.NewRandom(13))
	for tupleIndex := uint16(0); tupleIndex < 20; tupleIndex++ {
		tuple := table1.TUtils.RandTuple().Tuple
		err := page.InsertTuple(tuple.Bytes())
		if err != nil {
			t.Errorf("could not insert tuple nr %d because of error: %e", tupleIndex, err)
		}
		assert.Eq(page.TupleCount(), tupleIndex+1, t)
		assert.EqArray(tuple.Bytes(), page.Tuple(tupleIndex).Bytes(), t)
	}
}

func TestPage_DeleteTuple_FromMiddle(t *testing.T) {
	table1 := ttable1.Def(t)
	table1.TUtils.SetRand(random.NewRandom(23))
	tablePage := tpage.EmptyTablePage(table1, t)

	tuples := table1.TUtils.PutRandomTupleToPage(10, tablePage)
	tablePage.DeleteTuple(2)
	tablePage.DeleteTuple(8)

	assertTuplesList := []tpage.TupleIndex{0, 1, 3, 4, 5, 6, 7, 9}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(tablePage.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
}

func TestPage_DeleteTuple_First(t *testing.T) {
	table1 := ttable1.Def(t)
	table1.TUtils.SetRand(random.NewRandom(23))
	tablePage := tpage.EmptyTablePage(table1, t)

	tuples := table1.TUtils.PutRandomTupleToPage(10, tablePage)

	tablePage.DeleteTuple(0)
	assertTuplesList := []tpage.TupleIndex{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(tablePage.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
	tablePage.DeleteTuple(1)
	assertTuplesList = []tpage.TupleIndex{2, 3, 4, 5, 6, 7, 8, 9}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(tablePage.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
}

func TestPage_DeleteTuple_Last(t *testing.T) {
	table1 := ttable1.Def(t)
	table1.TUtils.SetRand(random.NewRandom(23))
	tablePage := tpage.EmptyTablePage(table1, t)

	tuples := table1.TUtils.PutRandomTupleToPage(10, tablePage)

	tablePage.DeleteTuple(9)
	assertTuplesList := []tpage.TupleIndex{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(tablePage.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
	tablePage.DeleteTuple(8)
	assertTuplesList = []tpage.TupleIndex{0, 1, 2, 3, 4, 5, 6, 7}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(tablePage.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
}
