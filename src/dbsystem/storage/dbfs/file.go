package dbfs

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"time"
)

type FileLike interface {
	io.WriterAt
	io.Writer
	io.ReaderAt
	io.Reader
	Name() string
	Stat() (fs.FileInfo, error)
	Close() error
}

var _ FileLike = &os.File{}

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
}

func (i *InMemoryFile) WriteAt(p []byte, off int64) (n int, err error) {
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
	n = copy(p, i.buffer[off:])
	if n != len(p) {
		return n, errors.New("n not equal to len(p)")
	}
	return n, nil
}

func (i *InMemoryFile) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryFile) Name() string {
	return i.stat.name
}

func (i *InMemoryFile) Stat() (fs.FileInfo, error) {
	return i.stat, nil
}

func (i *InMemoryFile) Close() error {
	i.buffer = nil
	i.stat = nil
	return nil
}

func (i *InMemoryFile) updateStat() {
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
