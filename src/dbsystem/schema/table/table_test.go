package table_test

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/ctypes"
	"HomegrownDB/dbsystem/schema/column/factory"
	"HomegrownDB/dbsystem/schema/table"
	"testing"
)

type TestTableBuilder struct{}

func (b TestTableBuilder) TestTable1() table.Definition {
	tableDef := table.NewDefinition("test_table1")
	tableDef.SetTableId(43741)
	tableDef.SetObjectId(642)
	tableDef.AddColumn(factory.CreateDefinition(column.ArgsBuilder("col1", ctypes.Int2).Build()))

	return tableDef
}

func TestTable_Serialization(t *testing.T) {

}
