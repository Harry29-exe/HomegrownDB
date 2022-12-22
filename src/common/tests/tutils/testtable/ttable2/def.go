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
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/schema/column"
	"testing"
)

func Def(t *testing.T) testtable.TestTable {
	table := testtable.NewTestTableBuilder(TableName).
		AddColumn(C0Id, false, hgtype.TypeInt8, hgtype.Args{}).
		AddColumn(C1Name, true, hgtype.TypeStr, hgtype.Args{}).
		AddColumn(C2Specie, true, hgtype.TypeStr, hgtype.Args{}).
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
