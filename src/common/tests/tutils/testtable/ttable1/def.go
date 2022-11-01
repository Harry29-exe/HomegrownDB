/*
Package ttable1 is following table definition

	CREATE TABLE awesome_table1 (

	awesome_key   int8  PRIMARY KEY,
	nullable_col  int8,
	non_null_coll  int8  NOT NULL

	);
*/
package ttable1

import (
	"HomegrownDB/common/tests/tutils/testtable"
	"HomegrownDB/dbsystem/ctype"
	"testing"
)

/*
Def creates following table definition

	CREATE TABLE awesome_table1 (
			awesome_key   int8  PRIMARY KEY,
			nullable_col  int8,
			non_null_coll  int8  NOT NULL
	);
*/
func Def(t *testing.T) testtable.TestTable {
	table := testtable.NewTestTableBuilder(TableName).
		AddColumn(C0AwesomeKey, false, ctype.TypeInt8, nil).
		AddColumn(C1NullableCol, true, ctype.TypeInt8, nil).
		AddColumn(C2NonNullColl, false, ctype.TypeInt8, nil).
		GetTable()

	return testtable.NewTestTable(table, t)
}
