package tutils

import (
	"HomegrownDB/common/tests/tstructs"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/ctypes"
)

var TestTables = testTables{}

type testTables struct{}

/*
Table1 creates following table definition
	CREATE TABLE awesome_table1 (
			awesome_key   int2  PRIMARY KEY,
			nullable_col  int2,
			non_null_coll  int2  NOT NULL
	);
*/
var Table1 = struct {
	AwesomeKey, NullableCol, NonNullColl string
}{"awesome_key", "nullable_col", "non_null_coll"}

func (t testTables) Table1Def() tstructs.TestTable {
	table := NewTestTableBuilder("awesome_table1").
		AddColumn(column.ArgsBuilder(Table1.AwesomeKey, ctypes.Int2).Build()).
		AddColumn(column.ArgsBuilder(Table1.NullableCol, ctypes.Int2).Build()).
		AddColumn(column.ArgsBuilder(Table1.NonNullColl, ctypes.Int2).Build()).
		GetTable()

	return tstructs.TestTable{WDefinition: table}
}
