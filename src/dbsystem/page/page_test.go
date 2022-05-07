package page_test

import (
	"HomegrownDB/dbsystem/page"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/testutils"
	"fmt"
	"testing"
)

func TestCreateEmptyPage(t *testing.T) {
	tableDef := table.NewDefinition(
		"test_table", 0, 0)

	newPage := page.CreateEmptyPage(tableDef)

	if newPage.TupleCount() != 0 {
		errMsg := fmt.Sprintf("new page has tuple count different than 0,\n page: %#v", newPage)
		t.Error(errMsg)
	}

	testutils.ShouldPanic(
		func() {
			newPage.Tuple(0)
		},
		"Newly created tuple returned non existing tuple with index 0", t)

	testutils.ShouldPanic(
		func() {
			newPage.UpdateTuple(0, make([]byte, 128))
		},
		"Newly created tuple updated non existing tuple with index 0", t)
}

func TestPage_Tuple(t *testing.T) {
	testPage := page.CreateEmptyPage()

}

var pUtils pageUtils = pageUtils{}

type pageUtils struct{}
