package hgtest

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/storage/dbfs"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

type TestFS interface {
	dbfs.FS
	Destroy()
}

func TestRootPath(t *testing.T) string {
	now := time.Now().String()
	index := strings.Index(now, ".")
	return fmt.Sprintf("/tmp/HomegrownDB_TEST-%s-%s", t.Name(), now[:index])
}

func CreateAndInitTestFS(t *testing.T) TestFS {
	now := time.Now().String()
	index := strings.Index(now, ".")
	rootPath := fmt.Sprintf("/tmp/HomegrownDB_TEST-%s-%s", t.Name(), now[:index])

	innerFS, err := dbfs.CreateFS(rootPath)
	if err != nil {
		t.Error(err.Error())
	}
	fs := &testFS{
		FS:       innerFS,
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
