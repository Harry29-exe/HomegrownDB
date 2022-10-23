package dbfs

import (
	"errors"
	"io/fs"
	"os"
	"sync"
	"sync/atomic"
)

var _ FileLike = &InMemLockableFile{}

func NewInMemLockableFile(filename string) *InMemLockableFile {
	return &InMemLockableFile{
		buffer: make([]byte, 0, 1000),
		stat: &fileInfo{
			name: filename,
			size: 0,
		},
		closed: false,

		lock:     &sync.RWMutex{},
		waitingR: 0,
		waitingW: 0,
	}
}

type InMemLockableFile struct {
	buffer []byte
	stat   *fileInfo
	closed bool

	lock     *sync.RWMutex
	waitingR int32
	waitingW int32
}

func (i *InMemLockableFile) WriteAt(p []byte, off int64) (n int, err error) {
	if i.closed {
		return 0, os.ErrClosed
	}
	i.beforeWrite()
	defer i.afterWrite()

	diff := len(i.buffer[off:]) - len(p)
	if diff >= 0 {
		return copy(i.buffer[off:], p), nil
	}

	copied := copy(i.buffer[off:], p)
	i.buffer = append(i.buffer, p[copied:]...)

	i.updateStat()
	return len(p), nil
}

func (i *InMemLockableFile) Write(p []byte) (n int, err error) {
	if i.closed {
		return 0, errors.New("file is closed")
	}
	i.beforeWrite()
	defer i.afterWrite()

	i.buffer = append(i.buffer, p...)

	i.updateStat()
	return len(p), nil
}

// todo compare to os.File implementation
func (i *InMemLockableFile) ReadAt(p []byte, off int64) (n int, err error) {
	if i.closed {
		return 0, os.ErrClosed
	}
	i.beforeRead()
	defer i.afterRead()

	n = copy(p, i.buffer[off:])
	if n != len(p) {
		return n, errors.New("n not equal to len(p)")
	}
	return n, nil
}

func (i *InMemLockableFile) Read(p []byte) (n int, err error) {
	if i.closed {
		return 0, os.ErrClosed
	}
	i.beforeRead()
	i.afterRead()

	//TODO implement me
	panic("implement me")
}

func (i *InMemLockableFile) Name() string {
	return i.stat.name
}

func (i *InMemLockableFile) Stat() (fs.FileInfo, error) {
	if i.closed {
		return nil, os.ErrClosed
	}
	return i.stat, nil
}

func (i *InMemLockableFile) Close() error {
	if i.closed {
		return os.ErrClosed
	}
	i.closed = true
	return nil
}

func (i *InMemLockableFile) Reopen() error {
	if !i.closed {
		return errors.New("file is not closed")
	}
	i.closed = false
	return nil
}

func (i *InMemLockableFile) Lock() {
	i.lock.Lock()
}

func (i *InMemLockableFile) Unlock() {
	i.lock.Unlock()
}

func (i *InMemLockableFile) RLock() {
	i.lock.RLock()
}

func (i *InMemLockableFile) RUnlock() {
	i.lock.RUnlock()
}

func (i *InMemLockableFile) GetReadWaiting() int32 {
	return i.waitingR
}

func (i *InMemLockableFile) GetWriteWaiting() int32 {
	return i.waitingW
}

func (i *InMemLockableFile) updateStat() {
	i.stat.size = int64(len(i.buffer))
}

func (i *InMemLockableFile) beforeWrite() {
	atomic.AddInt32(&i.waitingW, 1)
	i.lock.Lock()
	atomic.AddInt32(&i.waitingW, -1)
}

func (i *InMemLockableFile) afterWrite() {
	i.lock.Unlock()
}

func (i *InMemLockableFile) beforeRead() {
	atomic.AddInt32(&i.waitingR, 1)
	i.lock.RLock()
	atomic.AddInt32(&i.waitingR, -1)
}

func (i *InMemLockableFile) afterRead() {
	i.lock.RUnlock()
}
