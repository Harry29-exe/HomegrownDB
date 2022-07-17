package buffer_test

import (
	"HomegrownDB/common/tests/tstructs"
	"HomegrownDB/common/tests/tutils"
	"HomegrownDB/dbsystem/io/buffer"
	"testing"
)

var testTables = tutils.TestTables

func TestSharedBuffer_Overflow(t *testing.T) {
	tables := []tstructs.TestTable{
		tutils.TestTables.Table1Def(),
	}
	tableStore := tstructs.NewTestTableStoreWithInMemoryIO(tables)
	for i := 0; i < 10_000; i++ {

	}

	testBuffer := buffer.NewSharedBuffer(100, tableStore)

	if testBuffer != nil {

	}
	//testBuffer.RPage()
}
