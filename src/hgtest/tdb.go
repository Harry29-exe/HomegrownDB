package hgtest

import (
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/pageio"
	"testing"
)

func NewTestDB(db hg.DB, t *testing.T) TestDB {
	return TestDB{DB: db, T: t}
}

type TestDB struct {
	hg.DB
	T *testing.T
}

func (db TestDB) RandFillPages(pageToFill int, tableName string) {

}

func (db TestDB) TableByName(tableName string) table.Definition {
	id := db.TableStore().FindTable(tableName)
	if id == relation.InvalidRelId {
		db.T.Errorf("not table: " + tableName)
	}
	return db.TableStore().AccessTable(id, table.WLockMode)
}

func (db TestDB) PageIOByTableName(tableName string) pageio.IO {
	id := db.TableStore().FindTable(tableName)
	if id == relation.InvalidRelId {
		db.T.Errorf("not table: " + tableName)
	}
	return db.PageIOStore().Get(id)
}
