package hgtest

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/storage/dbfs"
	"os"
	"testing"
)

type TestFS interface {
	dbfs.FS
	Destroy()
}

func CreateAndInitTestFS(t *testing.T) TestFS {
	rootPath := "/tmp/HomegrownDB_TEST-*"
	err := os.Mkdir(rootPath, 0777)
	//rootPath, err := os.MkdirTemp("", "HomegrownDB_TEST-*")
	if err != nil {
		t.Error(err.Error())
	}

	fs := &testFS{
		FS:       dbfs.LoadFS(rootPath),
		Filepath: rootPath,
	}
	err = fs.InitDBSystem()
	assert.ErrIsNil(err, t)

	return fs
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

// -------------------------
//      FSUtils
// -------------------------

var FSUtils = fsUtils{}

type fsUtils struct{}
