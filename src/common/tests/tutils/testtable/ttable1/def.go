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
		AddColumn(C0AwesomeKey, false, C2NonNullCollType).
		AddColumn(C1NullableCol, true, C2NonNullCollType).
		AddColumn(C2NonNullColl, false, C2NonNullCollType).
		GetTable()

	return testtable.NewTestTable(table, t)
}
