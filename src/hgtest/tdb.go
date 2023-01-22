package hgtest

import (
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/storage/tpage"
	"testing"
)

func NewTestDBUtils(db hg.DB, t *testing.T) TestDBUtils {
	return TestDBUtils{
		DB:   db,
		Rand: random.NewRandom(1),
		T:    t,
	}
}

type TestDBUtils struct {
	DB   hg.DB
	Rand random.Random
	T    *testing.T
}

// -------------------------
//      Table Utils
// -------------------------

func (u TestDBUtils) FillTablePages(pagesToFill int, tableName string) {
	table, tableIO := u.TableByName(tableName), u.PageIOByTableName(tableName)
	filledPages := 0
	tablePage := tpage.AsPage(make([]byte, page.Size), page.Id(filledPages), table)
	insertedTuples := 0
	for filledPages < pagesToFill {
		err := tablePage.InsertTuple(u.RandTuple(table).Bytes())
		insertedTuples++

		if err != nil {
			err = tableIO.FlushPage(page.Id(filledPages), tablePage.Bytes())
			filledPages++
			if err != nil {
				panic("could not create new page: " + err.Error())
			}
			tablePage = tpage.AsPage(make([]byte, page.Size), page.Id(filledPages), table)
		}
	}
}

func (u TestDBUtils) TableByName(tableName string) table.Definition {
	id := u.DB.TableStore().FindTable(tableName)
	if id == relation.InvalidRelId {
		u.T.Errorf("not table: " + tableName)
	}
	return u.DB.TableStore().AccessTable(id, table.WLockMode)
}

func (u TestDBUtils) RandTuple(tableRel table.Definition) tpage.Tuple {
	return Table.RandTPageTuple(tableRel, u.Rand)
}

// -------------------------
//      PageIO Utils
// -------------------------

func (u TestDBUtils) PageIOByTableName(tableName string) pageio.IO {
	id := u.DB.TableStore().FindTable(tableName)
	if id == relation.InvalidRelId {
		u.T.Errorf("not table: " + tableName)
	}
	return u.DB.PageIOStore().Get(id)
}
