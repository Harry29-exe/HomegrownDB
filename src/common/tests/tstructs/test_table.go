package tstructs

import (
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
	"fmt"
)

type TestTable struct {
	table.WDefinition
}

func (t TestTable) FillPages(pagesToFill int, tableIO pageio.IO, rand random.Random) {
	tablePage := page.NewPage(t.WDefinition, make([]byte, pageSize))
	filledPages := 0
	insertedTuples := 0
	for filledPages < pagesToFill {
		err := tablePage.InsertTuple(t.RandTuple(rand).Tuple.Bytes())
		insertedTuples++

		if err != nil {
			filledPages++
			_, err = tableIO.NewPage(tablePage.Data())
			if err != nil {
				panic("could not create new page")
			}
			tablePage = page.NewPage(t.WDefinition, make([]byte, pageSize))
		}
	}
}

func (t TestTable) PutRandomTupleToPage(tupleCount int, tablePage page.TablePage, rand random.Random) []page.Tuple {
	tuples := make([]page.Tuple, tupleCount)
	for i := 0; i < tupleCount; i++ {
		tuple := t.RandTuple(rand).Tuple
		tuples[i] = tuple
		err := tablePage.InsertTuple(tuple.Bytes())
		if err != nil {
			panic(fmt.Sprintf("TestTable.PutRandomTupleToPage got error: %s", err.Error()))
		}
	}

	return tuples
}

func (t TestTable) RandTuple(rand random.Random) page.TupleToSave {
	values := map[string][]byte{}
	for i := uint16(0); i < t.ColumnCount(); i++ {
		col := t.Column(i)
		values[col.Name()] = col.CType().Rand(rand)
	}

	tuple, err := page.NewTestTuple(t.WDefinition, values, tx.NewInfoCtx(rand.Int31()))
	if err != nil {
		panic(err.Error())
	}

	return tuple
}
