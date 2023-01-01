package hgtest

import (
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/storage/tpage"
	"HomegrownDB/dbsystem/tx"
	"testing"
)

func NewTestDB(db hg.DB, t *testing.T) TestDBUtils {
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

func (u TestDBUtils) TestFillTablePages(pagesToFill int, tableName string) {
	table, tableIO := u.TableByName(tableName), u.PageIOByTableName(tableName)
	filledPages := 0
	tablePage := tpage.AsPage(make([]byte, page.Size), page.Id(filledPages), table)
	insertedTuples := 0
	for filledPages < pagesToFill {
		err := tablePage.InsertTuple(u.RandTuple(table).Tuple.Bytes())
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

func (u TestDBUtils) PageIOByTableName(tableName string) pageio.IO {
	id := u.DB.TableStore().FindTable(tableName)
	if id == relation.InvalidRelId {
		u.T.Errorf("not table: " + tableName)
	}
	return u.DB.PageIOStore().Get(id)
}

// -------------------------
//      internal
// -------------------------

func (u TestDBUtils) RandTuple(table table.Definition) tpage.TupleToSave {
	values := map[string][]byte{}
	for i := uint16(0); i < table.ColumnCount(); i++ {
		col := table.Column(i)
		values[col.Name()] = col.CType().Rand(u.Rand)
	}

	tuple, err := tpage.NewTestTuple(table, values, tx.NewInfoCtx(u.Rand.Int31()))
	if err != nil {
		panic(err.Error())
	}

	return tuple
}
