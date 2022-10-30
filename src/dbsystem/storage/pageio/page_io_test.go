package pageio_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"github.com/spf13/afero"
	"testing"
)

func TestPageIO_Reopen(t *testing.T) {
	inMemFile := dbfs.NewInMemoryFile("page_io")
	pageIO, err := pageio.NewPageIO(inMemFile)
	assert.IsNil(err, t)

	page0 := createSimplePage(0)
	err = pageIO.FlushPage(0, page0)
	assert.IsNil(err, t)
	err = pageIO.Close()
	assert.IsNil(err, t)

	assert.IsNil(inMemFile.Reopen(), t)

	buff := buffer()
	loadedPageIO, err := pageio.LoadPageIO(inMemFile)
	assert.IsNil(err, t)
	err = loadedPageIO.ReadPage(0, buff)
	assert.IsNil(err, t)
	assert.EqArray(page0, buff, t)
}

func TestFlushNewPageWithExistingPages(t *testing.T) {
	//given
	inMemFS := afero.NewMemMapFs()
	file, err := inMemFS.Create("/tmp/test.page")
	assert.IsNil(err, t)

	io, err := pageio.NewPageIO(file)
	assert.IsNil(err, t)

	pages := [][]byte{createSimplePage(0), createSimplePage(1), createSimplePage(2)}
	for i, newPage := range pages {
		err = io.FlushPage(page.Id(i), newPage)
		assert.IsNil(err, t)
	}

	//when
	page4 := createSimplePage(4)
	err = io.FlushPage(4, page4)
	assert.IsNil(err, t)

	//then
	for i := 0; i < 3; i++ {
		readPageAndAssertEq(i, pages[i], io, t)
	}
	emptyPage := make([]byte, page.Size)
	readPageAndAssertEq(3, emptyPage, io, t)
	readPageAndAssertEq(4, page4, io, t)
}

func TestFlushNewPageToEmptyIO(t *testing.T) {
	//given
	inMemFS := afero.NewMemMapFs()
	file, err := inMemFS.Create("/tmp/test.page")
	assert.IsNil(err, t)

	io, err := pageio.NewPageIO(file)
	assert.IsNil(err, t)

	//when
	page4 := createSimplePage(4)
	err = io.FlushPage(4, page4)
	assert.IsNil(err, t)

	//then
	emptyPage := make([]byte, page.Size)
	for i := 0; i < 4; i++ {
		readPageAndAssertEq(i, emptyPage, io, t)
	}
	readPageAndAssertEq(4, page4, io, t)
}

func readPageAndAssertEq(i int, expectedPage []byte, io pageio.IO, t *testing.T) {
	buff := make([]byte, page.Size)
	err := io.ReadPage(page.Id(i), buff)
	assert.IsNil(err, t)
	assert.EqArray(buff, expectedPage, t)
}

func createSimplePage(seqStart byte) []byte {
	data := make([]byte, page.Size)
	for i := 0; i < int(page.Size); i++ {
		data[i] = byte(i) + seqStart
	}
	return data
}

func buffer() []byte {
	return make([]byte, page.Size)
}
