package hgtest

import (
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/lib/random"
)

var Table = tableUtils{}

type tableUtils struct{}

func (t tableUtils) RandTPageTuple(table reldef.TableRDefinition, rand random.Random) page.Tuple {
	values := make([][]byte, table.ColumnCount())
	for i := uint16(0); i < table.ColumnCount(); i++ {
		col := table.Column(i)
		values[i] = col.CType().Rand(rand)
	}

	tuple := page.NewTuple(values, page.PatternFromTable(table), &tx.StdTx{Id: tx.Id(rand.Int31())})

	return tuple
}

func (t tableUtils) RandTuple(table reldef.TableRDefinition, rand random.Random) page.Tuple {
	values := make([][]byte, table.ColumnCount())
	for i := uint16(0); i < table.ColumnCount(); i++ {
		col := table.Column(i)
		values[i] = col.CType().Rand(rand)
	}

	tuple := page.NewTuple(values, page.PatternFromTable(table), &tx.StdTx{Id: tx.Id(rand.Int31())})

	return tuple
}
