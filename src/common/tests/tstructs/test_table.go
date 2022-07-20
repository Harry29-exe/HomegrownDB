package tstructs

import (
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/io"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/ctypes"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"fmt"
)

type TestTable struct {
	table.WDefinition
}

func (t TestTable) FillPages(pagesToFill int, tableIO io.TableDataIO, rand random.Random) {
	page := bdata.NewPage(t.WDefinition, make([]byte, pageSize))
	filledPages := 0
	insertedTuples := 0
	for filledPages < pagesToFill {
		err := page.InsertTuple(t.RandTuple(rand).Tuple.Data())
		insertedTuples++
		println(insertedTuples)
		if insertedTuples == 356 {
			print("about to crash")
		}

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

func (t TestTable) RandTuple(rand random.Random) bdata.TupleToSave {
	values := map[string]any{}
	for i := uint16(0); i < t.ColumnCount(); i++ {
		col := t.GetColumn(i)
		values[col.Name()] = t.randValueForColumnType(col.Type(), rand)
	}

	tuple, err := bdata.CreateTuple(t.WDefinition, values, tx.NewContext(rand.Int31()))
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
