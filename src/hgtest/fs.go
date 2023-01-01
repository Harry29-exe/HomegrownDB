package hgtest

import (
	"HomegrownDB/dbsystem/storage/dbfs"
	"os"
	"testing"
)

type TestFS interface {
	dbfs.FS
	Destroy()
}

func CreateNewTestFS(t *testing.T) dbfs.FS {
	rootPath, err := os.MkdirTemp("", "HomegrownDB_TEST-*")
	if err != nil {
		t.Error(err.Error())
	}

	return &testFS{
		FS:       dbfs.NewFS(rootPath),
		Filepath: rootPath,
	}
}

type testFS struct {
	dbfs.FS
	Filepath string
}

func (t *testFS) Destroy() {
	err := os.RemoveAll(t.Filepath)
	if err != nil {
		panic("could not remove file: " + t.Filepath)
	}
}
