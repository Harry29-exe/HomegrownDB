package testinfr

import (
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	table2 "HomegrownDB/dbsystem/relation/table"
	"testing"
)

var TestTableStore = tblStore{}

type tblStore struct{}

func (tblStore) TableStore(t *testing.T, tables ...table2.Definition) table2.Store {
	store := table2.NewEmptyTableStore()
	for _, tab := range tables {
		err := store.AddNewTable(tab)
		if err != nil {
			t.Error(err.Error())
		}
	}
	return store
}

func (tblStore) StoreWithUsersTable(t *testing.T) (store table2.Store, users table2.Definition) {
	store = table2.NewEmptyTableStore()
	users = tt_user.Def(t)
	users.InitRel(1, 2, 3)
	err := store.AddNewTable(users)
	if err != nil {
		t.Error(err.Error())
	}
	return store, users
}
