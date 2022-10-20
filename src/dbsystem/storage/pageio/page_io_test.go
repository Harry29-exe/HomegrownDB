package pageio_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/pageio"
	"testing"
)

func TestPageIO_Reopen(t *testing.T) {
	inMemFile := dbfs.NewInMemoryFile("page_io")
	pageIO := pageio.NewPageIO(inMemFile)

	page0 := createSimplePage()
	_, err := pageIO.NewPage(page0)
	assert.IsNil(err, t)
	err = pageIO.Close()
	assert.IsNil(err, t)

	assert.IsNil(inMemFile.Reopen(), t)

	buff := buffer()
	loadedPageIO := pageio.LoadPageIO(inMemFile)
	err = loadedPageIO.ReadPage(0, buff)
	assert.IsNil(err, t)
	assert.EqArray(page0, buff, t)
}

func TestPageIO_Locks(t *testing.T) {

}

func createSimplePage() []byte {
	data := make([]byte, dbbs.PageSize)
	for i := 0; i < int(dbbs.PageSize); i++ {
		data[i] = byte(i)
	}
	return data
}

func buffer() []byte {
	return make([]byte, dbbs.PageSize)
}
