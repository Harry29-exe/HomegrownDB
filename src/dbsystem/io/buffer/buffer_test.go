package buffer_test

import (
	"HomegrownDB/common/random"
	"HomegrownDB/common/tests"
	"HomegrownDB/common/tests/tstructs"
	"HomegrownDB/common/tests/tutils"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/io/buffer"
	"sync"
	"testing"
	"time"
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

func TestSharedBuffer_ParallelRead(t *testing.T) {
	table1 := tutils.TestTables.Table1Def()
	tableStore := tstructs.NewTestTableStoreWithInMemoryIO(table1)
	table1IO := tableStore.TableIO(table1.TableId())

	rand := random.NewRandom(0)
	table1.FillPages(10, table1IO, rand)

	testBuffer := buffer.NewSharedBuffer(10, tableStore)

	tCount := 4
	waitGroup1 := sync.WaitGroup{}
	waitGroup1.Add(tCount)
	waitGroup2 := sync.WaitGroup{}
	waitGroup2.Add(tCount)

	tag := bdata.PageTag{PageId: 0, TableId: table1.TableId()}
	for i := 0; i < tCount; i++ {
		go func() {
			_, _ = testBuffer.RPage(tag)
			waitGroup1.Done()
			waitGroup1.Wait()
			testBuffer.ReleaseRPage(tag)
			waitGroup2.Done()
		}()
	}

	waitGroup2.Wait()
}

//todo consult chaber if this test can be done better (not using timer)
func TestSharedBuffer_RWLock(t *testing.T) {
	table1 := tutils.TestTables.Table1Def()
	tableStore := tstructs.NewTestTableStoreWithInMemoryIO(table1)
	table1IO := tableStore.TableIO(table1.TableId())

	rand := random.NewRandom(0)
	table1.FillPages(10, table1IO, rand)

	testBuffer := buffer.NewSharedBuffer(10, tableStore)

	tag := bdata.NewPageTag(0, table1)
	_, err := testBuffer.WPage(tag)

	if err != nil {
		t.Fail()
		return
	}

	ch1 := make(chan bool)
	go func() {
		_, err := testBuffer.RPage(tag)
		ch1 <- true
		if err != nil {
			t.Errorf("testBuffer.RPage returned error: %e", err)
			return
		}
		testBuffer.ReleaseRPage(tag)
	}()

	ch2 := make(chan bool)
	go func() {
		time.Sleep(10 * time.Nanosecond)
		ch2 <- true
	}()

	select {
	case <-ch1:
		t.Errorf("goroutine was able to access rpage before time ran out")
		t.FailNow()
	case <-ch2:
		testBuffer.ReleaseWPage(tag)
	}

	<-ch1
}

//todo consult chaber if this test can be done better (not using timer)
func TestSharedBuffer_2xWLock(t *testing.T) {
	table1 := tutils.TestTables.Table1Def()
	tableStore := tstructs.NewTestTableStoreWithInMemoryIO(table1)
	table1IO := tableStore.TableIO(table1.TableId())

	rand := random.NewRandom(0)
	table1.FillPages(10, table1IO, rand)

	testBuffer := buffer.NewSharedBuffer(10, tableStore)

	tag := bdata.NewPageTag(0, table1)
	_, err := testBuffer.WPage(tag)

	if err != nil {
		t.Fail()
		return
	}

	ch1 := make(chan bool)
	go func() {
		_, err := testBuffer.WPage(tag)
		ch1 <- true
		if err != nil {
			t.Errorf("testBuffer.RPage returned error: %e", err)
			return
		}
		testBuffer.ReleaseWPage(tag)
	}()

	ch2 := make(chan bool)
	go func() {
		time.Sleep(10 * time.Nanosecond)
		ch2 <- true
	}()

	select {
	case <-ch1:
		t.Errorf("goroutine was able to access rpage before time ran out")
		t.FailNow()
	case <-ch2:
		testBuffer.ReleaseWPage(tag)
	}

	<-ch1
}
