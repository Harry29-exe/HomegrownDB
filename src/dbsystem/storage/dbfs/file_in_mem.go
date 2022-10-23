package dbfs

import (
	"errors"
	"io/fs"
	"os"
)

var _ FileLike = &InMemoryFile{}

func NewInMemoryFile(filename string) *InMemoryFile {
	return &InMemoryFile{
		buffer: make([]byte, 0, 1000),
		stat: &fileInfo{
			name: filename,
			size: 0,
		},
	}
}

type InMemoryFile struct {
	buffer []byte
	stat   *fileInfo
	closed bool
}

func (i *InMemoryFile) WriteAt(p []byte, off int64) (n int, err error) {
	if i.closed {
		return 0, os.ErrClosed
	}
	diff := len(i.buffer[off:]) - len(p)
	if diff >= 0 {
		return copy(i.buffer[off:], p), nil
	}

	copied := copy(i.buffer[off:], p)
	i.buffer = append(i.buffer, p[copied:]...)

	i.updateStat()
	return len(p), nil
}

func (i *InMemoryFile) Write(p []byte) (n int, err error) {
	i.buffer = append(i.buffer, p...)

	i.updateStat()
	return len(p), nil
}

// todo compare to os.File implementation
func (i *InMemoryFile) ReadAt(p []byte, off int64) (n int, err error) {
	if i.closed {
		return 0, os.ErrClosed
	}
	n = copy(p, i.buffer[off:])
	if n != len(p) {
		return n, errors.New("n not equal to len(p)")
	}
	return n, nil
}

func (i *InMemoryFile) Read(p []byte) (n int, err error) {
	if i.closed {
		return 0, os.ErrClosed
	}
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryFile) Name() string {
	return i.stat.name
}

func (i *InMemoryFile) Stat() (fs.FileInfo, error) {
	if i.closed {
		return nil, os.ErrClosed
	}
	return i.stat, nil
}

func (i *InMemoryFile) Close() error {
	if i.closed {
		return os.ErrClosed
	}
	i.closed = true
	return nil
}

func (i *InMemoryFile) Reopen() error {
	if !i.closed {
		return errors.New("file is not closed")
	}
	i.closed = false
	return nil
}

func (i *InMemoryFile) updateStat() {
	i.stat.size = int64(len(i.buffer))
}
