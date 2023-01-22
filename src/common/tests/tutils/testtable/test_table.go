package testtable

import (
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/storage/tpage"
	"HomegrownDB/dbsystem/tx"
	"fmt"
	"testing"
)

func NewTestTable(def table.Definition, t *testing.T) TestTable {
	return TestTable{
		Definition: def,
		TUtils: TUtils{
			table: def,
			rand:  random.NewRandom(0),
		},
		T: t,
	}
}

type TestTable struct {
	table.Definition
	TUtils TUtils
	T      *testing.T
}

var pageSize = config.PageSize

type TUtils struct {
	rand  random.Random
	table table.Definition
}

func (t *TUtils) SetRand(rand random.Random) {
	t.rand = rand
}

func (t *TUtils) FillPages(pagesToFill int, tableIO pageio.IO) {
	filledPages := 0
	tablePage := tpage.AsPage(make([]byte, pageSize), page.Id(filledPages), t.table)
	insertedTuples := 0
	for filledPages < pagesToFill {
		err := tablePage.InsertTuple(t.RandTuple().Bytes())
		insertedTuples++

		if err != nil {
			err = tableIO.FlushPage(page.Id(filledPages), tablePage.Bytes())
			filledPages++
			if err != nil {
				panic("could not create new page")
			}
			tablePage = tpage.AsPage(make([]byte, pageSize), page.Id(filledPages), t.table)
		}
	}
}

func (t *TUtils) PutRandomTupleToPage(tupleCount int, tablePage tpage.Page) []tpage.Tuple {
	tuples := make([]tpage.Tuple, tupleCount)
	for i := 0; i < tupleCount; i++ {
		tuple := t.RandTuple()
		tuples[i] = tuple
		err := tablePage.InsertTuple(tuple.Bytes())
		if err != nil {
			panic(fmt.Sprintf("TestTable.PutRandomTupleToPage got error: %s", err.Error()))
		}
	}

	return tuples
}

func (t *TUtils) RandTuple() tpage.Tuple {
	values := make([][]byte, t.table.ColumnCount())
	for i := uint16(0); i < t.table.ColumnCount(); i++ {
		col := t.table.Column(i)
		values[i] = col.CType().Rand(t.rand)
	}

	tuple := tpage.NewTuple(values, t.table, &tx.StdTx{Id: tx.Id(t.rand.Int31())})

	return tuple
}
