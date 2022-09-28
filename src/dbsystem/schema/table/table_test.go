package table_test

import (
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"testing"
)

type TestTableBuilder struct{}

func (b TestTableBuilder) TestTable1() table.Definition {
	tableDef := table.NewDefinition("test_table1")
	tableDef.SetTableId(43741)
	tableDef.SetObjectId(642)
	err := tableDef.AddColumn(column.NewDefinition(column.ArgsBuilder("col1", ctype.TypeInt8).Build()))
	if err != nil {
		panic("unexpected error")
	}

	return tableDef
}

func TestTable_Serialization(t *testing.T) {

}
