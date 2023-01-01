package buffer_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/common/tests/tutils/testtable/ttable1"
	. "HomegrownDB/dbsystem/access/buffer"
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
	var stdBuffer StdBuffer
	buffSize, pageToRead := uint(100), 1_000

	ctx := bootstrapUsersTCtx(func(args di.SimpleArgs, store pageio.Store) (SharedBuffer, error) {
		stdBuffer = NewStdBuffer(buffSize, store)
		return AsSharedBuffer(stdBuffer), nil
	}, t)

	ctx.TestFillTablePages(pageToRead, tt_user.TableName)

	rawPage := make([]byte, page.Size)
	// when
	for i := page.Id(0); i < 1_000; i++ {
		tag := pageio.NewTablePageTag(i, ctx.table)
		pageData, err := stdBuffer.ReadWPage(ctx.table, i, RbmReadOrCreate)
		assert.ErrIsNil(err, t)

		err = ctx.tableIO.ReadPage(i, rawPage)
		assert.ErrIsNil(err, t)
		stdBuffer.ReleaseWPage(tag)

		// then
		assert.EqArray(pageData.Bytes, rawPage, t)
	}
}

func TestSharedBuffer_ParallelRead(t *testing.T) {
	//given
	tb := table1Bootstrap(t)
	table, tableIO, ioStore := tb.table, tb.tableIO, tb.ioStore

	tCount := 16
	table.TUtils.FillPages(10, tableIO)

	testBuffer := NewStdBuffer(2, ioStore)

	//when
	waitGroup1 := sync.WaitGroup{}
	waitGroup1.Add(tCount)
	waitGroup2 := sync.WaitGroup{}
	waitGroup2.Add(tCount)

	tag := pageio.PageTag{PageId: 0, Relation: table.RelationID()}
	for i := 0; i < tCount; i++ {
		go func() {
			_, _ = testBuffer.ReadRPage(table, 0, RbmReadOrCreate)
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

	testBuffer := NewStdBuffer(uint(tCount), ioStore)

	waitGroup1 := sync.WaitGroup{}
	waitGroup1.Add(tCount)
	lock := &sync.RWMutex{}
	waitGroup2 := sync.WaitGroup{}
	waitGroup2.Add(tCount)

	for i := 0; i < tCount; i++ {
		pageId := page.Id(i)
		go func() {
			_, _ = testBuffer.ReadRPage(table1, pageId, RbmReadOrCreate)
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

	testBuffer := NewStdBuffer(10, ioStore)

	tag := pageio.NewTablePageTag(0, table1)
	_, err := testBuffer.ReadWPage(table1, 0, RbmReadOrCreate)

	if err != nil {
		t.Fail()
		return
	}

	ch1 := make(chan bool)
	go func() {
		_, err = testBuffer.ReadRPage(table1, 0, RbmReadOrCreate)
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

	testBuffer := NewStdBuffer(10, ioStore)

	tag := pageio.NewTablePageTag(0, table1)
	_, err := testBuffer.ReadWPage(table1, 0, RbmReadOrCreate)

	if err != nil {
		t.Fail()
		return
	}

	ch1 := make(chan bool)
	go func() {
		_, err = testBuffer.ReadWPage(table1, 0, RbmReadOrCreate)
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
//      Bootstrap
// -------------------------

type usersTCtx struct {
	hgtest.TestDBUtils

	table table.Definition

	tableIO pageio.IO
}

func bootstrapUsersTCtx(provider di.SharedBufferProvider, t *testing.T) *usersTCtx {
	rootPath := hgtest.TestRootPath(t)
	err := hg.Create(hg.CreateArgs{Mode: hg.Test, RootPath: rootPath})
	assert.ErrIsNil(err, t)

	container := hg.DefaultFutureContainer()
	container.RootProvider = func() (string, error) { return rootPath, nil }
	container.SharedBufferProvider = provider
	db, err := hg.Load(&container)
	assert.ErrIsNil(err, t)
	testDB := hgtest.NewTestDB(db, t)

	usersTable := tt_user.Def(t)
	err = db.CreateRel(usersTable)
	assert.ErrIsNil(err, t)

	return &usersTCtx{
		TestDBUtils: testDB,
		table:       usersTable,
		tableIO:     testDB.PageIOByTableName(tt_user.TableName),
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
