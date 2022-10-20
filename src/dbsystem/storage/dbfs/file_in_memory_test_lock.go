package dbfs

import (
	"HomegrownDB/common/datastructs/appsync"
	"errors"
	"io/fs"
	"os"
	"sync"
	"time"
)

var _ FileLike = &InMemoryFileWithTestLock{}

func NewInMemoryFileWithLocks(filename string) *InMemoryFileWithTestLock {
	return &InMemoryFileWithTestLock{
		buffer: make([]byte, 0, 1000),
		stat: &fileInfo{
			name: filename,
			size: 0,
		},
		closed: false,

		lock:     &sync.RWMutex{},
		waitingR: appsync.NewSyncCounter(0),
		waitingW: appsync.NewSyncCounter(0),
	}
}

type InMemoryFileWithTestLock struct {
	buffer []byte
	stat   *fileInfo
	closed bool

	lock     *sync.RWMutex
	waitingR appsync.SyncCounter[int]
	waitingW appsync.SyncCounter[int]
}

func (i *InMemoryFileWithTestLock) WriteAt(p []byte, off int64) (n int, err error) {
	if i.closed {
		return 0, os.ErrClosed
	}
	i.waitingW.IncrementAndGet()
	i.lock.Lock()

	diff := len(i.buffer[off:]) - len(p)
	if diff >= 0 {
		return copy(i.buffer[off:], p), nil
	}

	copied := copy(i.buffer[off:], p)
	i.buffer = append(i.buffer, p[copied:]...)

	i.updateStat()
	return len(p), nil
}

func (i *InMemoryFileWithTestLock) Write(p []byte) (n int, err error) {
	i.buffer = append(i.buffer, p...)

	i.updateStat()
	return len(p), nil
}

// todo compare to os.File implementation
func (i *InMemoryFileWithTestLock) ReadAt(p []byte, off int64) (n int, err error) {
	if i.closed {
		return 0, os.ErrClosed
	}
	n = copy(p, i.buffer[off:])
	if n != len(p) {
		return n, errors.New("n not equal to len(p)")
	}
	return n, nil
}

func (i *InMemoryFileWithTestLock) Read(p []byte) (n int, err error) {
	if i.closed {
		return 0, os.ErrClosed
	}
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryFileWithTestLock) Name() string {
	return i.stat.name
}

func (i *InMemoryFileWithTestLock) Stat() (fs.FileInfo, error) {
	if i.closed {
		return nil, os.ErrClosed
	}
	return i.stat, nil
}

func (i *InMemoryFileWithTestLock) Close() error {
	if i.closed {
		return os.ErrClosed
	}
	i.closed = true
	return nil
}

func (i *InMemoryFileWithTestLock) Reopen() error {
	if !i.closed {
		return errors.New("file is not closed")
	}
	i.closed = false
	return nil
}

func (i *InMemoryFileWithTestLock) updateStat() {
	i.stat.size = int64(len(i.buffer))
}

var _ os.FileInfo = &fileInfo{}

type fileInfo struct {
	name string
	size int64
}

func (f *fileInfo) Name() string {
	return f.name
}

func (f *fileInfo) Size() int64 {
	return f.size
}

func (f *fileInfo) Mode() fs.FileMode {
	//TODO implement me
	panic("implement me")
}

func (f *fileInfo) ModTime() time.Time {
	//TODO implement me
	panic("implement me")
}

func (f *fileInfo) IsDir() bool {
	return false
}

func (f *fileInfo) Sys() any {
	//TODO implement me
	panic("implement me")
}
