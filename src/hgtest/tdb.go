package hgtest

import (
	"HomegrownDB/common/random"
	table2 "HomegrownDB/dbsystem/access/reldef/tabdef"
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
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
	tablePage := page.AsTablePage(make([]byte, page.Size), page.Id(filledPages), table)
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
			tablePage = page.AsTablePage(make([]byte, page.Size), page.Id(filledPages), table)
		}
	}
}

func (u TestDBUtils) TableByName(tableName string) tabdef.Definition {
	id := u.DB.RelationManager().FindTable(tableName)
	if id == reldef.InvalidRelId {
		u.T.Errorf("not tabdef: " + tableName)
	}
	return u.DB.RelationManager().AccessTable(id, table2.WLockMode)
}

func (u TestDBUtils) RandTuple(tableRel tabdef.Definition) page.Tuple {
	return Table.RandTPageTuple(tableRel, u.Rand)
}

// -------------------------
//      PageIO Utils
// -------------------------

func (u TestDBUtils) PageIOByTableName(tableName string) pageio.IO {
	id := u.DB.RelationManager().FindTable(tableName)
	if id == reldef.InvalidRelId {
		u.T.Errorf("not tabdef: " + tableName)
	}
	return u.DB.PageIOStore().Get(id)
}
