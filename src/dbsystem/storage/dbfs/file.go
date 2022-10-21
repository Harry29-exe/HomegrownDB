package dbfs

import (
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
