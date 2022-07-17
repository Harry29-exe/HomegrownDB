package testutils

import (
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/schema/table"
	"math/rand"
)

type TestTable struct {
	table.WDefinition
}

func (t TestTable) RandTuple(random rand.Rand) bdata.Tuple {
	for i := uint16(0); i < t.ColumnCount(); i++ {
		col := t.GetColumn(i)

	}
}
