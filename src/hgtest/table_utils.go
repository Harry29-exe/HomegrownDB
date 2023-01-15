package hgtest

import (
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/tpage"
	"HomegrownDB/dbsystem/tx"
)

var Table = tableUtils{}

type tableUtils struct{}

func (t tableUtils) RandTuple(table table.Definition, rand random.Random) tpage.TupleToSave {
	values := map[string][]byte{}
	for i := uint16(0); i < table.ColumnCount(); i++ {
		col := table.Column(i)
		values[col.Name()] = col.CType().Rand(rand)
	}

	tuple, err := tpage.NewTestTuple(table, values, tx.NewInfoCtx(rand.Int31()))
	if err != nil {
		panic(err.Error())
	}

	return tuple
}
