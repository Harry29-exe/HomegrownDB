package buffer_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/ttable1"
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/storage/pageio"
	"github.com/spf13/afero"
	"testing"
)

func TestTableBufferWriteRead(t *testing.T) {
	//given
	pageioStore := pageio.NewStore()
	fs := afero.NewMemMapFs()

	table1 := ttable1.Def(t)
	file, err := fs.Create("table1_io")
	assert.IsNil(err, t)
	io, err := pageio.NewPageIO(file)
	assert.IsNil(err, t)
	pageioStore.Register(table1.RelationId(), io)

	buff := buffer.NewSharedBuffer(2, pageioStore)

	//when

	//inserting data
	wPage0, err := buff.WTablePage(0, table1)
	wPage0Copy := make([]byte, dbsystem.PageSize)
	copy(wPage0Copy, wPage0.Bytes())
	assert.IsNil(err, t)
	p0Tuple0 := table1.TUtils.RandTuple().Tuple.Bytes()
	err = wPage0.InsertTuple(p0Tuple0)
	assert.IsNil(err, t)

	wPage1, err := buff.WTablePage(1, table1)
	assert.IsNil(err, t)
	wPage1Copy := make([]byte, dbsystem.PageSize)
	copy(wPage1Copy, wPage1.Bytes())
	p1Tuple0 := table1.TUtils.RandTuple().Tuple.Bytes()
	err = wPage1.InsertTuple(p1Tuple0)
	assert.IsNil(err, t)

	buff.WPageRelease(wPage0.PageTag())
	buff.WPageRelease(wPage1.PageTag())

	//flushing page0 and page1 to 'disc'
	wPage2, err := buff.WTablePage(2, table1)
	wPage3, err := buff.WTablePage(3, table1)
	buff.WPageRelease(wPage2.PageTag())
	buff.WPageRelease(wPage3.PageTag())

	//then
	// page0
	wPage0, err = buff.WTablePage(0, table1)
	assert.IsNil(err, t)
	assert.EqArray(wPage0Copy, wPage0.Bytes(), t)

	wPage0FromIO := make([]byte, dbsystem.PageSize)
	err = io.ReadPage(0, wPage0FromIO)
	assert.IsNil(err, t)
	assert.EqArray(wPage0.Bytes(), wPage0FromIO, t)

	//page1
	wPage1, err = buff.WTablePage(1, table1)
	assert.IsNil(err, t)
	assert.EqArray(wPage1Copy, wPage1.Bytes(), t)

	wPage1FromIO := make([]byte, dbsystem.PageSize)
	err = io.ReadPage(1, wPage0FromIO)
	assert.IsNil(err, t)
	assert.EqArray(wPage1Copy, wPage1FromIO, t)
}
