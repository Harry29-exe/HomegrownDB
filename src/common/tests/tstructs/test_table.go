package tstructs

import (
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"fmt"
)

type TestTable struct {
	table.WDefinition
}

func (t TestTable) FillPages(pagesToFill int, tableIO access.TableDataIO, rand random.Random) {
	page := dbbs.NewPage(t.WDefinition, make([]byte, pageSize))
	filledPages := 0
	insertedTuples := 0
	for filledPages < pagesToFill {
		err := page.InsertTuple(t.RandTuple(rand).Tuple.Bytes())
		insertedTuples++

		if err != nil {
			filledPages++
			_, err = tableIO.NewPage(page.Data())
			if err != nil {
				panic("could not create new page")
			}
			page = dbbs.NewPage(t.WDefinition, make([]byte, pageSize))
		}
	}
}

func (t TestTable) PutRandomTupleToPage(tupleCount int, page dbbs.Page, rand random.Random) []dbbs.Tuple {
	tuples := make([]dbbs.Tuple, tupleCount)
	for i := 0; i < tupleCount; i++ {
		tuple := t.RandTuple(rand).Tuple
		tuples[i] = tuple
		err := page.InsertTuple(tuple.Bytes())
		if err != nil {
			panic(fmt.Sprintf("TestTable.PutRandomTupleToPage got error: %s", err.Error()))
		}
	}

	return tuples
}

func (t TestTable) RandTuple(rand random.Random) dbbs.TupleToSave {
	values := map[string][]byte{}
	for i := uint16(0); i < t.ColumnCount(); i++ {
		col := t.Column(i)
		values[col.Name()] = col.CType().Rand(rand)
	}

	tuple, err := dbbs.CreateTuple(t.WDefinition, values, tx.NewInfoCtx(rand.Int31()))
	if err != nil {
		panic(err.Error())
	}

	return tuple
}
