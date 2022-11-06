package buffer_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable"
	"HomegrownDB/common/tests/tutils/testtable/ttable1"
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/storage/tpage"
	"github.com/spf13/afero"
	"testing"
)

func TestTableBufferWriteRead(t *testing.T) {
	//given
	table1, io, buff := createTestSharedBuffer(t)

	//when

	//inserting data
	wPage0 := insertTPageWithSingleTuple(0, table1, buff, t)
	wPage0Copy := copyTPage(wPage0)
	wPage1 := insertTPageWithSingleTuple(1, table1, buff, t)
	wPage1Copy := copyTPage(wPage1)

	buff.WPageRelease(wPage0.PageTag())
	buff.WPageRelease(wPage1.PageTag())

	//flushing page0 and page1 to 'disc'
	wPage2, err := buff.WTablePage(table1, 2)
	assert.IsNil(err, t)
	wPage3, err := buff.WTablePage(table1, 3)
	assert.IsNil(err, t)
	buff.WPageRelease(wPage2.PageTag())
	buff.WPageRelease(wPage3.PageTag())

	//then
	checkIfPageIsSaved(0, wPage0Copy, table1, io, buff, t)
	checkIfPageIsSaved(1, wPage1Copy, table1, io, buff, t)
}

func checkIfPageIsSaved(pageId page.Id, expectedPage []byte, table1 testtable.TestTable, io pageio.IO, buff buffer.SharedBuffer, t *testing.T) {
	wPage, err := buff.WTablePage(table1, pageId)
	assert.IsNil(err, t)
	assert.EqArray(expectedPage, wPage.Bytes(), t)

	wPage0FromIO := make([]byte, dbsystem.PageSize)
	err = io.ReadPage(pageId, wPage0FromIO)
	assert.IsNil(err, t)
	assert.EqArray(wPage.Bytes(), wPage0FromIO, t)
}

func insertTPageWithSingleTuple(pageId page.Id, table1 testtable.TestTable, buff buffer.SharedBuffer, t *testing.T) tpage.TableWPage {
	wPage0, err := buff.WTablePage(table1, pageId)
	assert.IsNil(err, t)
	p0Tuple0 := table1.TUtils.RandTuple().Tuple.Bytes()
	err = wPage0.InsertTuple(p0Tuple0)
	assert.IsNil(err, t)

	return wPage0
}

func copyTPage(tPage tpage.TableRPage) (pageCopy []byte) {
	pageCopy = make([]byte, dbsystem.PageSize)
	switch tablePage := tPage.(type) {
	case tpage.TablePage:
		copy(pageCopy, tablePage.Bytes())
	default:
		panic("not known TableRPage type")
	}

	return
}

func createTestSharedBuffer(t *testing.T) (testtable.TestTable, pageio.IO, buffer.SharedBuffer) {
	pageioStore := pageio.NewStore()
	fs := afero.NewMemMapFs()

	table1 := ttable1.Def(t)
	file, err := fs.Create("table1_io")
	assert.IsNil(err, t)
	io, err := pageio.NewPageIO(file)
	assert.IsNil(err, t)
	pageioStore.Register(table1.RelationID(), io)

	buff := buffer.NewSharedBuffer(2, pageioStore)
	return table1, io, buff
}
