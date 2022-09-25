package tstructs

import (
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/ctypes"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"fmt"
)

type TestTable struct {
	table.WDefinition
}

func (t TestTable) FillPages(pagesToFill int, tableIO access.TableDataIO, rand random.Random) {
	page := bdata.NewPage(t.WDefinition, make([]byte, pageSize))
	filledPages := 0
	insertedTuples := 0
	for filledPages < pagesToFill {
		err := page.InsertTuple(t.RandTuple(rand).Tuple.Data())
		insertedTuples++

		if err != nil {
			filledPages++
			_, err := tableIO.NewPage(page.Data())
			if err != nil {
				panic("could not create new page")
			}
			page = bdata.NewPage(t.WDefinition, make([]byte, pageSize))
		}
	}
}

func (t TestTable) PutRandomTupleToPage(tupleCount int, page bdata.Page, rand random.Random) []bdata.Tuple {
	tuples := make([]bdata.Tuple, tupleCount)
	for i := 0; i < tupleCount; i++ {
		tuple := t.RandTuple(rand).Tuple
		tuples[i] = tuple
		err := page.InsertTuple(tuple.Data())
		if err != nil {
			panic(fmt.Sprintf("TestTable.PutRandomTupleToPage got error: %s", err.Error()))
		}
	}

	return tuples
}

func (t TestTable) RandTuple(rand random.Random) bdata.TupleToSave {
	values := map[string]any{}
	for i := uint16(0); i < t.ColumnCount(); i++ {
		col := t.GetColumn(i)
		values[col.Name()] = t.randValueForColumnType(col.Type(), rand)
	}

	tuple, err := bdata.CreateTuple(t.WDefinition, values, tx.NewInfoCtx(rand.Int31()))
	if err != nil {
		panic(err.Error())
	}

	return tuple
}

func (t TestTable) randValueForColumnType(ctype column.Type, random random.Random) any {
	switch ctype {
	case ctypes.Int2:
		return random.Int16()

	default:
		panic(fmt.Sprintf("type %s not implemented ad TestTable.randValueForColumnType",
			ctype))
	}
}
