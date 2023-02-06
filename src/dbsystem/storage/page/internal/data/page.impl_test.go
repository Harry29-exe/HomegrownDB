package data_test

import (
	"HomegrownDB/common/random"
	"HomegrownDB/common/tests"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/ttable1"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/storage/page/internal/data"
	"HomegrownDB/hgtest"
	"fmt"
	"testing"
)

func TestCreateEmptyPage(t *testing.T) {
	tableDef := tabdef.NewDefinition(
		"test_table")
	tableDef.SetOID(32)
	tableDef.SetOID(12)

	newPage := data.EmptyTablePage(data.PatternFromTable(tableDef), tableDef.OID())

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
	// when
	table1 := ttable1.Def(t)
	newPage := data.EmptyTablePage(data.PatternFromTable(table1), table1.OID())

	for tupleIndex := uint16(0); tupleIndex < 20; tupleIndex++ {
		tuple := table1.TUtils.RandTuple()
		err := newPage.InsertTuple(tuple.Bytes())
		if err != nil {
			t.Errorf("could not insert tuple nr %d because of error: %e", tupleIndex, err)
		}
		assert.Eq(newPage.TupleCount(), tupleIndex+1, t)
		assert.EqArray(tuple.Bytes(), newPage.Tuple(tupleIndex).Bytes(), t)
	}
}

func TestPage_DeleteTuple_FromMiddle(t *testing.T) {
	// given
	table1 := ttable1.Def(t)
	newPage := data.EmptyTablePage(data.PatternFromTable(table1), table1.OID())

	tuples := putRandomTupleToPage(10, newPage, table1, t)

	// when
	newPage.DeleteTuple(2)
	newPage.DeleteTuple(8)

	// then
	assertTuplesList := []data.TupleIndex{0, 1, 3, 4, 5, 6, 7, 9}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(newPage.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
}

func TestPage_DeleteTuple_First(t *testing.T) {
	// given
	table1 := ttable1.Def(t)
	newPage := data.EmptyTablePage(data.PatternFromTable(table1), table1.OID())

	tuples := putRandomTupleToPage(10, newPage, table1, t)

	// when
	newPage.DeleteTuple(0)
	// then
	assertTuplesList := []data.TupleIndex{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(newPage.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
	// when
	newPage.DeleteTuple(1)
	//then
	assertTuplesList = []data.TupleIndex{2, 3, 4, 5, 6, 7, 8, 9}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(newPage.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
}

func TestPage_DeleteTuple_Last(t *testing.T) {
	// given
	table1 := ttable1.Def(t)
	newPage := data.EmptyTablePage(data.PatternFromTable(table1), table1.OID())

	tuples := putRandomTupleToPage(10, newPage, table1, t)

	// when
	newPage.DeleteTuple(9)
	// then
	assertTuplesList := []data.TupleIndex{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(newPage.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
	// when
	newPage.DeleteTuple(8)
	// then
	assertTuplesList = []data.TupleIndex{0, 1, 2, 3, 4, 5, 6, 7}
	for _, tupleIndex := range assertTuplesList {
		assert.EqArray(newPage.Tuple(tupleIndex).Bytes(), tuples[tupleIndex].Bytes(), t)
	}
}

func putRandomTupleToPage(tupleCount int, page data.Page, table tabdef.RDefinition, t *testing.T) []data.Tuple {
	rand := random.NewRandom(0)
	tuples := make([]data.Tuple, tupleCount)
	for i := 0; i < tupleCount; i++ {
		tuple := hgtest.Table.RandTuple(table, rand)
		tuples[i] = tuple
		err := page.InsertTuple(tuple.Bytes())
		assert.ErrIsNil(err, t)
	}
	return tuples
}
