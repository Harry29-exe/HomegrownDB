package dbtable_test

import (
	. "HomegrownDB/sql/schema"
	. "HomegrownDB/sql/schema/dbtable"
	"testing"
)

func TestTableSerialization(t *testing.T) {

}

func createTestTable() *Table {
	tb := NewTableBuilder("test_table")
	tb.AddNewColumn("col1_in8", *GetColumnType(Int8, nil), false, false)
	tb.AddNewColumn("col2_int4", *GetColumnType(Int4, nil), false, false)
	tb.AddNewColumn("col3_str", *GetColumnType(Varchar, []int32{128}), false, false)

	return tb.Build()
}
