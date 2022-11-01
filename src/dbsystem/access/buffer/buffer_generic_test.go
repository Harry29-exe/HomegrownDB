package buffer

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable"
	"HomegrownDB/common/tests/tutils/testtable/ttable1"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"github.com/spf13/afero"
	"sync"
	"testing"
	"time"
)

func TestSharedBuffer_Overflow(t *testing.T) {
	table1 := ttable1.Def(t)
	ioStore := pageio.NewStore()
	table1IO := _createAndRegisterTestPageIO(table1, ioStore, t)

	table1.TUtils.FillPages(1_000, table1IO)

	buf := make([]byte, page.Size)
	testBuffer := newSharedBuffer(100, ioStore)
	for i := page.Id(0); i < 1_000; i++ {
		tag := NewTablePageTag(i, table1)
		pageData, err := testBuffer.WPage(tag)
		if err != nil {
			t.Errorf("During reading page %d got error: %e", i, err)
		}

		err = table1IO.ReadPage(i, buf)
		if err != nil {
			t.Errorf("TableIO returned non nil error: %e", err)
		}
		testBuffer.ReleaseWPage(tag)
		assert.EqArray(pageData.bytes, buf, t)
	}
}

func TestSharedBuffer_ParallelRead(t *testing.T) {
	table1 := ttable1.Def(t)
	ioStore := pageio.NewStore()
	table1IO := _createAndRegisterTestPageIO(table1, ioStore, t)

	table1.TUtils.FillPages(10, table1IO)

	testBuffer := newSharedBuffer(10, ioStore)

	tCount := 16
	waitGroup1 := sync.WaitGroup{}
	waitGroup1.Add(tCount)
	waitGroup2 := sync.WaitGroup{}
	waitGroup2.Add(tCount)

	tag := page.Tag{PageId: 0, Relation: table1.RelationId()}
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

// todo consult chaber if this test can be done better (not using timer)
func TestSharedBuffer_RWLock(t *testing.T) {
	table1 := ttable1.Def(t)
	ioStore := pageio.NewStore()
	table1IO := _createAndRegisterTestPageIO(table1, ioStore, t)

	table1.TUtils.FillPages(10, table1IO)

	testBuffer := newSharedBuffer(10, ioStore)

	tag := NewTablePageTag(0, table1)
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
			t.Errorf("testBuffer.RTablePage returned error: %e", err)
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

// todo consult chaber if this test can be done better (not using timer)
func TestSharedBuffer_2xWLock(t *testing.T) {
	table1 := ttable1.Def(t)
	ioStore := pageio.NewStore()
	table1IO := _createAndRegisterTestPageIO(table1, ioStore, t)

	table1.TUtils.FillPages(10, table1IO)

	testBuffer := newSharedBuffer(10, ioStore)

	tag := NewTablePageTag(0, table1)
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
			t.Errorf("testBuffer.RTablePage returned error: %e", err)
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

func _createAndRegisterTestPageIO(table1 testtable.TestTable, ioStore *pageio.Store, t *testing.T) pageio.IO {
	fs := afero.NewMemMapFs()
	file, err := fs.Create("table1IO")
	assert.IsNil(err, t)
	table1IO, err := pageio.NewPageIO(file)
	assert.IsNil(err, t)
	ioStore.Register(table1.RelationId(), table1IO)
	return table1IO
}
