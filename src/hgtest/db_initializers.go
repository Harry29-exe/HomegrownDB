package hgtest

import (
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/lib/tests/assert"
	"HomegrownDB/lib/tests/tutils/testtable/tt_user"
	"testing"
)

type TestDBBuilder struct {
	db hg.DB
	t  *testing.T
	tx tx.Tx

	funcToExec []func()
}

func CreateAndLoadDBWith(builders *hg.MBuilders, t *testing.T) *TestDBBuilder {
	rootPath := TestRootPath(t)
	err := hg.Create(hg.CreateArgs{Mode: hg.CreatorModeTest, RootPath: rootPath})
	assert.ErrIsNil(err, t)

	if builders == nil {
		builders = hg.DefaultMBuilders()
	}
	builders.StorageMBuilder.RootPathProvider = createRootPathProvider(rootPath)
	db, err := hg.Load(builders)
	assert.ErrIsNil(err, t)

	t.Cleanup(func() {
		if err := db.FS().DestroyDB(); err != nil {
			t.Error(err)
		}
	})

	return &TestDBBuilder{db: db, t: t, tx: tx.StdTx{}}
}

func (b *TestDBBuilder) WithUsersTable() *TestDBBuilder {
	usersTableFunc := func() {
		users := tt_user.Def(b.t)
		_, err := b.db.AccessModule().RelationManager().Create(users, b.tx)
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

func createRootPathProvider(rootPath string) func() (string, error) {
	return func() (string, error) {
		return rootPath, nil
	}
}
