package dbfs_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/page"
	"sync"
	"testing"
	"time"
)

func TestInMemoryFileWithTestLock_RLocks(t *testing.T) {
	file, data := createInMemoryFileWithTestLock(t)
	file.Lock()

	waitInitEnd, waitEnd := &sync.WaitGroup{}, &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waitEnd.Add(1)
		waitInitEnd.Add(1)
		go func() {
			buff := make([]byte, len(data))
			waitInitEnd.Done()
			_, err := file.ReadAt(buff, 0)
			assert.IsNil(err, t)
			assert.EqArray(data, buff, t)
			waitEnd.Done()
		}()
	}

	waitInitEnd.Wait()
	assertHasRWaiting(file, 10, t)

	file.Unlock()
	waitEnd.Wait()
}

func TestImMemoryFileWithTestLock_WLocks(t *testing.T) {
	file, data := createInMemoryFileWithTestLock(t)
	file.Lock()

	waitInitEnd, waitEnd := &sync.WaitGroup{}, &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waitEnd.Add(1)
		waitInitEnd.Add(1)
		go func() {
			buff := make([]byte, len(data))
			waitInitEnd.Done()
			_, err := file.Write(buff)
			assert.IsNil(err, t)
			waitEnd.Done()
		}()
	}

	waitInitEnd.Wait()
	assertHasWWaiting(file, 10, t)

	file.Unlock()
	waitEnd.Wait()
}

func assertHasRWaiting(file *dbfs.InMemLockableFile, rWaiting int32, t *testing.T) {
	for i := 0; i < 10; i++ {
		if file.GetReadWaiting() == rWaiting {
			break
		}
		time.Sleep(1 * time.Microsecond)
	}
	assert.Eq(rWaiting, file.GetReadWaiting(), t)
}

func assertHasWWaiting(file *dbfs.InMemLockableFile, wWaiting int32, t *testing.T) {
	for i := 0; i < 10; i++ {
		if file.GetWriteWaiting() == wWaiting {
			break
		}
		time.Sleep(1 * time.Microsecond)
	}
	assert.Eq(wWaiting, file.GetWriteWaiting(), t)
}

func createInMemoryFileWithTestLock(t *testing.T) (
	file *dbfs.InMemLockableFile, data []byte,
) {
	filename := "test_file_42"
	file = dbfs.NewInMemLockableFile(filename)
	data = make([]byte, page.Size)
	_, err := file.Write(data)
	assert.IsNil(err, t)
	return
}
