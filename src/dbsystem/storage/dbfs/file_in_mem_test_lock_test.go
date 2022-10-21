package dbfs_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/storage/dbfs"
	"sync"
	"testing"
	"time"
)

func TestInMemoryFileWithTestLock_RLocks(t *testing.T) {
	filename := "test_file_42"
	file := dbfs.NewInMemoryFileWithLocks(filename)
	data := make([]byte, dbbs.PageSize)
	_, err := file.Write(data)
	assert.IsNil(err, t)
	file.Lock()

	waitInitEnd := &sync.WaitGroup{}
	waitEnd := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waitEnd.Add(1)
		waitInitEnd.Add(1)
		go func() {
			buff := make([]byte, dbbs.PageSize)
			waitInitEnd.Done()
			_, err = file.ReadAt(buff, 0)
			assert.IsNil(err, t)
			assert.EqArray(data, buff, t)
			waitEnd.Done()
		}()
	}

	waitInitEnd.Wait()
	for i := 0; i < 10; i++ {
		if file.GetReadWaiting() == 10 {
			break
		}
		time.Sleep(1 * time.Microsecond)
	}
	assert.Eq(10, file.GetReadWaiting(), t)
	file.Unlock()
	waitEnd.Wait()
}
