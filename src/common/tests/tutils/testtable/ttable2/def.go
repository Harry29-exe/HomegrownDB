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
	"HomegrownDB/dbsystem/hgtype/coltype"
	"HomegrownDB/dbsystem/relation/table/column"
	"testing"
)

func Def(t *testing.T) testtable.TestTable {
	table := testtable.NewTestTableBuilder(TableName).
		AddColumn(C0Id, false, C0IdType).
		AddColumn(C1Name, true, C1NameType).
		AddColumn(C2Specie, true, C2SpecieType).
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

var (
	C0IdType     = coltype.NewInt8(hgtype.Args{})
	C1NameType   = coltype.NewStr(hgtype.Args{Length: 255, VarLen: true, UTF8: true})
	C2SpecieType = coltype.NewStr(hgtype.Args{Length: 255, VarLen: true, UTF8: true})
)
