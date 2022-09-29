package tutils

import (
	"HomegrownDB/common/tests/tstructs"
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
)

var TestTables = testTables{}

type testTables struct{}

/*
Table1 creates following table definition
	CREATE TABLE awesome_table1 (
			awesome_key   int8  PRIMARY KEY,
			nullable_col  int8,
			non_null_coll  int8  NOT NULL
	);
*/
var Table1 = struct {
	AwesomeKey, NullableCol, NonNullColl       string
	AwesomeKeyId, NullableColId, NonNullCollId column.Order
}{"awesome_key", "nullable_col", "non_null_coll",
	0, 1, 2}

func (t testTables) Table1Def() tstructs.TestTable {
	table := NewTestTableBuilder("awesome_table1").
		AddColumn(Table1.AwesomeKey, false, ctype.TypeInt8, nil).
		AddColumn(Table1.NullableCol, true, ctype.TypeInt8, nil).
		AddColumn(Table1.NonNullColl, false, ctype.TypeInt8, nil).
		GetTable()

	return tstructs.TestTable{WDefinition: table}
}

const (
	Table1Name                       = "awesome_table1"
	Table1AwesomeKey    string       = "awesome_key"
	Table1NullableCol   string       = "nullable_col"
	Table1NonNullColl   string       = "non_null_coll"
	Table1AwesomeKeyId  column.Order = 0
	Table1NullableColId column.Order = 1
	Table1NonNullCollId column.Order = 2
)
