package buffer_test

import (
	"HomegrownDB/common/tests/testmocks"
	"HomegrownDB/common/tests/testutils"
	"HomegrownDB/dbsystem/io/buffer"
	"HomegrownDB/dbsystem/schema/table"
	"testing"
)

var testTables = testutils.TestTables

func TestSharedBuffer_Overflow(t *testing.T) {
	tables := []table.WDefinition{
		testutils.TestTables.Table1Def(),
	}
	tableStore := testmocks.NewTestTableStoreWithInMemoryIO(tables)
	testBuffer := buffer.NewSharedBuffer(10_000, tableStore)

	if testBuffer != nil {

	}
	//testBuffer.RPage()
}
