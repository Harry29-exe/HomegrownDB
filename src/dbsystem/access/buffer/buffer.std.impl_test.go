package buffer_test

import (
	. "HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/hgtest"
	"HomegrownDB/lib/tests/assert"
	"HomegrownDB/lib/tests/tutils/testtable/tt_user"
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
		tag := page.NewTablePageTag(i, tableRel)
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
	threadCount := 32
	buffSize := threadCount
	ctx, testBuffer := bootstrapSimpleTCtx_Users(uint(buffSize), t)
	tableRel := ctx.table

	ctx.TestDBUtils.FillTablePages(100, tt_user.TableName)

	//when
	waitGroup1 := sync.WaitGroup{}
	waitGroup1.Add(threadCount)
	waitGroup2 := sync.WaitGroup{}
	waitGroup2.Add(1)
	waitGroup3 := sync.WaitGroup{}
	waitGroup3.Add(threadCount)

	tag := page.PageTag{PageId: 0, OwnerID: tableRel.OID()}
	for i := 0; i < threadCount; i++ {
		go func() {
			_, _ = testBuffer.ReadRPage(tableRel.OID(), 0, RbmReadOrCreate)
			waitGroup1.Done()
			waitGroup1.Wait()
			waitGroup2.Wait()
			testBuffer.ReleaseRPage(tag)
			waitGroup3.Done()
		}()
	}
	// wait for all goroutines to be inited
	waitGroup1.Wait()
	// lock threadCount - 1 pages (we want to make sure we have threadCount-1 free buffer pages)
	for i := 1; i < threadCount; i++ {
		_, err := testBuffer.ReadRPage(tableRel.OID(), page.Id(i), RbmReadOrCreate)
		assert.ErrIsNil(err, t)
	}
	// free locked pages
	for i := 1; i < threadCount; i++ {
		_, err := testBuffer.ReadRPage(tableRel.OID(), page.Id(i), RbmReadOrCreate)
		assert.ErrIsNil(err, t)
	}
	waitGroup2.Done()
	//then
	waitGroup3.Wait()
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
			tag := page.PageTag{PageId: pageId, OwnerID: tableRel.OID()}
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

	tag := page.NewTablePageTag(0, tableRel)
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

	tag := page.NewTablePageTag(0, usersTable)
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

	table reldef.TableDefinition

	tableIO pageio.IO
}

func bootstrapTCtx_Users(provider func(storageModule storage.Module, configModule config.Module) (SharedBuffer, error), t *testing.T) *usersTCtx {
	rootPath := hgtest.TestRootPath(t)
	err := hg.CreateDB(hg.CreateArgs{Mode: hg.InstallerModeTest, RootPath: rootPath})
	assert.ErrIsNil(err, t)

	container := hg.DefaultMBuilders()
	container.StorageMBuilder.RootPathProvider = func() (string, error) { return rootPath, nil }
	container.AccessMBuilder.SharedBufferProvider = provider
	db, err := hg.Load(container)
	assert.ErrIsNil(err, t)
	testDB := hgtest.NewTestDBUtils(db, t)

	usersTable := tt_user.Def(t)
	_, err = db.AccessModule().RelationManager().Create(usersTable, tx.StdTx{})
	assert.ErrIsNil(err, t)

	return &usersTCtx{
		TestDBUtils: testDB,
		table:       usersTable,
		tableIO:     testDB.PageIOByTableName(tt_user.TableName),
	}
}

func bootstrapSimpleTCtx_Users(bufferSize uint, t *testing.T) (*usersTCtx, StdBuffer) {
	var testBuffer StdBuffer
	ctx := bootstrapTCtx_Users(func(storageModule storage.Module, configModule config.Module) (SharedBuffer, error) {
		testBuffer = NewStdBuffer(bufferSize, storageModule.PageIOStore())
		return AsSharedBuffer(testBuffer), nil
	}, t)
	return ctx, testBuffer
}
