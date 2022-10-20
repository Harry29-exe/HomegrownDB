package dbfs

import (
	"io"
	"io/fs"
	"os"
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
