/*
Package ttable2 is following table definition

	CREATE TABLE birds (
		id int8 PRIMARY KEY,
		name varchar(255),
		specie varchar(255)
	)
*/
package ttable2

import (
	"HomegrownDB/common/tests/tutils/testtable"
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
	"testing"
)

func Def(t *testing.T) testtable.TestTable {
	table := testtable.NewTestTableBuilder(TableName).
		AddColumn(C0Id, false, ctype.TypeInt8, nil).
		AddColumn(C1Name, true, ctype.TypeStr, nil).
		AddColumn(C2Specie, true, ctype.TypeStr, nil).
		GetTable()

	return testtable.NewTestTable(table, t)
}

const (
	TableName                  = "birds"
	C0Id          string       = "id"
	C0IdOrder     column.Order = 0
	C1Name        string       = "name"
	C1NameOrder   column.Order = 1
	C2Specie      string       = "specie"
	C2SpecieOrder column.Order = 2
)
