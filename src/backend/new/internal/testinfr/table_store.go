package testinfr

import (
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/schema/table"
	"testing"
)

var TestTableStore = tblStore{}

type tblStore struct{}

func (tblStore) TableStore(t *testing.T, tables ...table.Definition) table.Store {
	store := table.NewEmptyTableStore()
	for _, tab := range tables {
		err := store.AddNewTable(tab)
		if err != nil {
			t.Error(err.Error())
		}
	}
	return store
}

func (tblStore) StoreWithUsersTable(t *testing.T) (store table.Store, users table.Definition) {
	store = table.NewEmptyTableStore()
	users = tt_user.Def(t)
	err := store.AddNewTable(users)
	if err != nil {
		t.Error(err.Error())
	}
	return store, users
}
