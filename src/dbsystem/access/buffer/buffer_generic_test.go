package buffer

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/common/tests/tutils/testtable/ttable1"
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/dbsystem/hg/di"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/hgtest"
	"sync"
	"testing"
	"time"
)

func TestSharedBuffer_Overflow(t *testing.T) {
	// given
	var internalBuff internalBuffer
	ctx := bootstrap(func(args di.SimpleArgs, store pageio.Store) (SharedBuffer, error) {
		internalBuff = newBuffer(100, store)
		return &bufferProxy{buffer: internalBuff}, nil
	}, t)
	ctx.db.RandFillPages(1_000, ctx.usersTable.Name())

	// when
	tmpPage := make([]byte, page.Size)
	for i := page.Id(0); i < 1_000; i++ {
		tag := pageio.NewTablePageTag(i, ctx.usersTable)
		pageData, err := internalBuff.ReadWPage(ctx.usersTable, i, rbmReadOrCreate)
		assert.ErrIsNil(err, t)

		err = ctx.usersIO.ReadPage(i, tmpPage)
		assert.ErrIsNil(err, t)
		internalBuff.ReleaseWPage(tag)

		// then
		assert.EqArray(pageData.bytes, tmpPage, t)
	}
}

func TestSharedBuffer_ParallelRead(t *testing.T) {
	//given
	tb := table1Bootstrap(t)
	table, tableIO, ioStore := tb.table, tb.tableIO, tb.ioStore

	tCount := 16
	table.TUtils.FillPages(10, tableIO)

	testBuffer := newBuffer(2, ioStore)

	//when
	waitGroup1 := sync.WaitGroup{}
	waitGroup1.Add(tCount)
	waitGroup2 := sync.WaitGroup{}
	waitGroup2.Add(tCount)

	tag := pageio.PageTag{PageId: 0, Relation: table.RelationID()}
	for i := 0; i < tCount; i++ {
		go func() {
			_, _ = testBuffer.ReadRPage(table, 0, rbmReadOrCreate)
			waitGroup1.Done()
			waitGroup1.Wait()
			testBuffer.ReleaseRPage(tag)
			waitGroup2.Done()
		}()
	}

	//then
	waitGroup2.Wait()
}

func TestSharedBuffer_ParallelDifferentRowRead(t *testing.T) {
	tb := table1Bootstrap(t)
	table1, table1IO, ioStore := tb.table, tb.tableIO, tb.ioStore

	tCount := 16
	table1.TUtils.FillPages(tCount, table1IO)

	testBuffer := newBuffer(uint(tCount), ioStore)

	waitGroup1 := sync.WaitGroup{}
	waitGroup1.Add(tCount)
	lock := &sync.RWMutex{}
	waitGroup2 := sync.WaitGroup{}
	waitGroup2.Add(tCount)

	for i := 0; i < tCount; i++ {
		pageId := page.Id(i)
		go func() {
			_, _ = testBuffer.ReadRPage(table1, pageId, rbmReadOrCreate)
			waitGroup1.Done()
			waitGroup1.Wait()
			lock.RLock()
			tag := pageio.PageTag{PageId: pageId, Relation: table1.RelationID()}
			testBuffer.ReleaseRPage(tag)
			waitGroup2.Done()
		}()
	}

	lock.Lock()
	waitGroup1.Wait()
	lock.Unlock()
	waitGroup2.Wait()
}

// todo consult chaber if this test can be done better (not using timer)
func TestSharedBuffer_RWLock(t *testing.T) {
	tb := table1Bootstrap(t)
	table1, table1IO, ioStore := tb.table, tb.tableIO, tb.ioStore

	table1.TUtils.FillPages(10, table1IO)

	testBuffer := newBuffer(10, ioStore)

	tag := pageio.NewTablePageTag(0, table1)
	_, err := testBuffer.ReadWPage(table1, 0, rbmReadOrCreate)

	if err != nil {
		t.Fail()
		return
	}

	ch1 := make(chan bool)
	go func() {
		_, err = testBuffer.ReadRPage(table1, 0, rbmReadOrCreate)
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
	tb := table1Bootstrap(t)
	table1, table1IO, ioStore := tb.table, tb.tableIO, tb.ioStore

	table1.TUtils.FillPages(10, table1IO)

	testBuffer := newBuffer(10, ioStore)

	tag := pageio.NewTablePageTag(0, table1)
	_, err := testBuffer.ReadWPage(table1, 0, rbmReadOrCreate)

	if err != nil {
		t.Fail()
		return
	}

	ch1 := make(chan bool)
	go func() {
		_, err = testBuffer.ReadWPage(table1, 0, rbmReadOrCreate)
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

// -------------------------
//      Bootstrap new
// -------------------------

type ctx struct {
	t  *testing.T
	db hgtest.TestDB

	usersTable table.Definition
	usersIO    pageio.IO
}

func bootstrap(bufferProvider di.SharedBufferProvider, t *testing.T) ctx {
	err := hg.Create(hg.CreateArgs{Mode: hg.Test, RootPath: hgtest.TestRootPath(t)})
	assert.ErrIsNil(err, t)

	fc := hg.DefaultFutureContainer()
	fc.SharedBufferProvider = bufferProvider
	db, err := hg.Load(&fc)
	tdb := hgtest.TestDB{DB: db, T: t}
	assert.ErrIsNil(err, t)

	assert.ErrIsNil(db.CreateRel(tt_user.Def(t)), t)

	return ctx{
		t:          t,
		db:         tdb,
		usersTable: tdb.TableByName(tt_user.TableName),
		usersIO:    tdb.PageIOByTableName(tt_user.TableName),
	}
}

// -------------------------
//      Bootstrap
// -------------------------

// genericBuffTB bootstrap for generic buffer tests
type genericBuffTB struct {
	t *testing.T

	table   testtable.TestTable
	tableIO pageio.IO

	fs      hgtest.TestFS
	ioStore pageio.Store
}

func (b genericBuffTB) CleanUp() {
	b.fs.Destroy()
}

func table1Bootstrap(t *testing.T) genericBuffTB {
	table1 := ttable1.Def(t)
	fs := hgtest.CreateAndInitTestFS(t)
	ioStore := hgtest.PageIOUtils.With(t, fs, table1)
	table1IO := ioStore.Get(table1.RelationID())

	return genericBuffTB{
		t:       t,
		table:   table1,
		tableIO: table1IO,
		fs:      fs,
		ioStore: ioStore,
	}
}
