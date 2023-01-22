package buffer_test

import (
	"HomegrownDB/common/random"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/common/tests/tutils/testtable/ttable1"
	. "HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/dbsystem/hg/di"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/hgtest"
	"testing"
)

func TestTableBufferWriteRead(t *testing.T) {
	//given
	fc := hg.DefaultFutureContainer()
	fc.SharedBufferProvider = func(args di.SimpleArgs, store pageio.Store) (SharedBuffer, error) {
		return NewSharedBuffer(2, store), nil
	}
	dbUtils := hgtest.CreateAndLoadDBWith(&fc, t).
		WithUsersTable().
		Build()
	table1, tableIO, buff := dbUtils.TableByName(tt_user.TableName),
		dbUtils.PageIOByTableName(tt_user.TableName),
		dbUtils.DB.SharedBuffer()

	//when

	//inserting data
	wPage0 := insertTPageWithSingleTuple(0, table1, buff, dbUtils.Rand, t)
	wPage0Copy := copyTPage(wPage0)
	wPage1 := insertTPageWithSingleTuple(1, table1, buff, dbUtils.Rand, t)
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
	checkIfPageIsSaved(0, wPage0Copy, table1, tableIO, buff, t)
	checkIfPageIsSaved(1, wPage1Copy, table1, tableIO, buff, t)
}

func checkIfPageIsSaved(pageId page.Id, expectedPage []byte, table1 table.Definition, io pageio.IO, buff SharedBuffer, t *testing.T) {
	wPage, err := buff.WTablePage(table1, pageId)
	assert.IsNil(err, t)
	assert.EqArray(expectedPage, wPage.Bytes(), t)

	wPage0FromIO := make([]byte, config.PageSize)
	err = io.ReadPage(pageId, wPage0FromIO)
	assert.IsNil(err, t)
	assert.EqArray(wPage.Bytes(), wPage0FromIO, t)
}

func insertTPageWithSingleTuple(pageId page.Id, table1 table.Definition, buff SharedBuffer, rand random.Random, t *testing.T) page.WPage {
	wPage0, err := buff.WTablePage(table1, pageId)
	assert.IsNil(err, t)
	p0Tuple0 := hgtest.Table.RandTPageTuple(table1, rand).Bytes()
	err = wPage0.InsertTuple(p0Tuple0)
	assert.IsNil(err, t)

	return wPage0
}

func copyTPage(tPage page.RPage) (pageCopy []byte) {
	pageCopy = make([]byte, config.PageSize)
	tPage.CopyBytes(pageCopy)
	return
}

func createTestSharedBuffer(t *testing.T) (testtable.TestTable, pageio.IO, SharedBuffer) {
	table1 := ttable1.Def(t)

	fs := hgtest.CreateAndInitTestFS(t)
	pageioStore := hgtest.PageIOUtils.With(t, fs, table1)

	buff := NewSharedBuffer(2, pageioStore)
	return table1, pageioStore.Get(table1.OID()), buff
}
