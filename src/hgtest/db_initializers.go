package hgtest

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/dbsystem/hg/di"
	"testing"
)

type TestDBBuilder struct {
	db hg.DB
	t  *testing.T

	funcToExec []func()
}

func CreateAndLoadDBWith(fc *di.FutureContainer, t *testing.T) *TestDBBuilder {
	rootPath := TestRootPath(t)
	err := hg.Create(hg.CreateArgs{Mode: hg.CreatorModeTest, RootPath: rootPath})
	assert.ErrIsNil(err, t)

	if fc == nil {
		temp := hg.DefaultFutureContainer()
		fc = &temp
	}
	fc.RootProvider = createRootPathProvider(rootPath)
	db, err := hg.Load(fc)
	assert.ErrIsNil(err, t)

	t.Cleanup(func() {
		if err := db.FS().DestroyDB(); err != nil {
			t.Error(err)
		}
	})

	return &TestDBBuilder{db: db, t: t}
}

func (b *TestDBBuilder) WithUsersTable() *TestDBBuilder {
	usersTableFunc := func() {
		users := tt_user.Def(b.t)
		err := b.db.CreateRel(users)
		assert.ErrIsNil(err, b.t)
	}

	b.funcToExec = append(b.funcToExec, usersTableFunc)
	return b
}

func (b *TestDBBuilder) Build() TestDBUtils {
	for _, fn := range b.funcToExec {
		fn()
	}

	b.t.Cleanup(func() {
		if err := b.db.Destroy(); err != nil {
			b.t.Error(err)
		}
	})
	return NewTestDBUtils(b.db, b.t)
}

// -------------------------
//      test providers
// -------------------------

func createRootPathProvider(rootPath string) di.RootPathProvider {
	return func() (string, error) {
		return rootPath, nil
	}
}
