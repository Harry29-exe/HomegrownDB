package buffer_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
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
	buffSize, pageToRead := uint(100), 1_000
	ctx, testBuffer := bootstrapSimpleTCtx_Users(buffSize, t)
	tableRel := ctx.table

	ctx.FillTablePages(pageToRead, tt_user.TableName)
	rawPage := make([]byte, page.Size)

	// when
	for i := page.Id(0); i < 1_000; i++ {
		tag := pageio.NewTablePageTag(i, tableRel)
		pageData, err := testBuffer.ReadWPage(tableRel.OID(), i, RbmReadOrCreate)
		assert.ErrIsNil(err, t)

		err = ctx.tableIO.ReadPage(i, rawPage)
		assert.ErrIsNil(err, t)
		testBuffer.ReleaseWPage(tag)

		// then
		assert.EqArray(pageData.Bytes, rawPage, t)
	}
}

func TestSharedBuffer_ParallelRead(t *testing.T) {
	//given
	buffSize := 2
	ctx, testBuffer := bootstrapSimpleTCtx_Users(uint(buffSize), t)
	tableRel := ctx.table

	tCount := 32
	ctx.TestDBUtils.FillTablePages(10, tt_user.TableName)

	//when
	waitGroup1 := sync.WaitGroup{}
	waitGroup1.Add(tCount)
	waitGroup2 := sync.WaitGroup{}
	waitGroup2.Add(tCount)

	tag := pageio.PageTag{PageId: 0, OwnerID: tableRel.OID()}
	for i := 0; i < tCount; i++ {
		go func() {
			_, _ = testBuffer.ReadRPage(tableRel.OID(), 0, RbmReadOrCreate)
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
	//given
	tCount := 32
	buffSize := tCount

	ctx, testBuffer := bootstrapSimpleTCtx_Users(uint(buffSize), t)
	tableRel := ctx.table
	ctx.TestDBUtils.FillTablePages(tCount, tableRel.Name())

	// when
	waitGroup1 := sync.WaitGroup{}
	waitGroup1.Add(tCount)
	lock := &sync.RWMutex{}
	waitGroup2 := sync.WaitGroup{}
	waitGroup2.Add(tCount)

	for i := 0; i < tCount; i++ {
		pageId := page.Id(i)
		go func() {
			_, _ = testBuffer.ReadRPage(tableRel.OID(), pageId, RbmReadOrCreate)
			waitGroup1.Done()
			waitGroup1.Wait()
			lock.RLock()
			tag := pageio.PageTag{PageId: pageId, OwnerID: tableRel.OID()}
			testBuffer.ReleaseRPage(tag)
			waitGroup2.Done()
		}()
	}

	lock.Lock()
	waitGroup1.Wait()
	lock.Unlock()
	// then
	waitGroup2.Wait()
}

// todo consult chaber if this test can be done better (not using timer)
func TestSharedBuffer_RWLock(t *testing.T) {
	ctx, testBuffer := bootstrapSimpleTCtx_Users(10, t)
	tableRel := ctx.table
	ctx.FillTablePages(10, tableRel.Name())

	tag := pageio.NewTablePageTag(0, tableRel)
	_, err := testBuffer.ReadWPage(tableRel.OID(), 0, RbmReadOrCreate)
	assert.ErrIsNil(err, t)

	ch1 := make(chan bool)
	go func() {
		_, err = testBuffer.ReadRPage(tableRel.OID(), 0, RbmReadOrCreate)
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
	ctx, testBuffer := bootstrapSimpleTCtx_Users(10, t)
	usersTable := ctx.table

	ctx.FillTablePages(10, tt_user.TableName)

	tag := pageio.NewTablePageTag(0, usersTable)
	_, err := testBuffer.ReadWPage(usersTable.OID(), 0, RbmReadOrCreate)

	if err != nil {
		t.Fail()
		return
	}

	ch1 := make(chan bool)
	go func() {
		_, err = testBuffer.ReadWPage(usersTable.OID(), 0, RbmReadOrCreate)
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
	testBuffer StdBuffer

	table table.Definition

	tableIO pageio.IO
}

func bootstrapTCtx_Users(provider di.SharedBufferProvider, t *testing.T) *usersTCtx {
	rootPath := hgtest.TestRootPath(t)
	err := hg.Create(hg.CreateArgs{Mode: hg.Test, RootPath: rootPath})
	assert.ErrIsNil(err, t)

	container := hg.DefaultFutureContainer()
	container.RootProvider = func() (string, error) { return rootPath, nil }
	container.SharedBufferProvider = provider
	db, err := hg.Load(&container)
	assert.ErrIsNil(err, t)
	testDB := hgtest.NewTestDBUtils(db, t)

	usersTable := tt_user.Def(t)
	err = db.CreateRel(usersTable)
	assert.ErrIsNil(err, t)

	return &usersTCtx{
		TestDBUtils: testDB,
		table:       usersTable,
		tableIO:     testDB.PageIOByTableName(tt_user.TableName),
	}
}

func bootstrapSimpleTCtx_Users(bufferSize uint, t *testing.T) (*usersTCtx, StdBuffer) {
	var testBuffer StdBuffer
	ctx := bootstrapTCtx_Users(func(args di.SimpleArgs, store pageio.Store) (SharedBuffer, error) {
		testBuffer = NewStdBuffer(bufferSize, store)
		return AsSharedBuffer(testBuffer), nil
	}, t)
	return ctx, testBuffer
}
