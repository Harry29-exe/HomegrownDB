package buffer_test

import (
	"HomegrownDB/common/random"
	"HomegrownDB/common/tests"
	"HomegrownDB/common/tests/tstructs"
	"HomegrownDB/common/tests/tutils"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/io/buffer"
	"testing"
)

func TestSharedBuffer_Overflow(t *testing.T) {
	table1 := tutils.TestTables.Table1Def()
	tableStore := tstructs.NewTestTableStoreWithInMemoryIO(table1)
	table1IO := tableStore.TableIO(table1.TableId())

	rand := random.NewRandom(0)
	table1.FillPages(1_000, table1IO, rand)

	buf := make([]byte, bdata.PageSize)
	testBuffer := buffer.NewSharedBuffer(100, tableStore)
	for i := bdata.PageId(0); i < 1_000; i++ {
		tag := bdata.NewPageTag(i, table1)
		page, err := testBuffer.WPage(tag)
		if err != nil {
			t.Errorf("During reading page %d got error: %e", i, err)
		}

		err = table1IO.ReadPage(i, buf)
		if err != nil {
			t.Errorf("TableIO returned non nil error: %e", err)
		}
		testBuffer.ReleaseWPage(tag)
		data := page.Data()
		tests.AssertEqArray(data, buf, t)
	}
}
