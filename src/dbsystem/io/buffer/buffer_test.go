package buffer_test

import (
	"HomegrownDB/dbsystem/io/buffer"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/tests/testmocks"
	"HomegrownDB/tests/testutils"
	"testing"
)

var testTables = testutils.TestTables

func TestSharedBuffer_Overflow(t *testing.T) {
	tables := []table.WDefinition{
		testutils.TestTables.Table1Def(),
	}
	tableStore := testmocks.NewTestTableStoreWithInMemoryIO(tables)
	testBuffer := buffer.NewSharedBuffer(10_000, tableStore)

	testBuffer.RPage()
}
