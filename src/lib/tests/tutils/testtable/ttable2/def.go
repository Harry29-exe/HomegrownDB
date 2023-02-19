/*
Package ttable2 is following tabdef definition

	CREATE TABLE birds (
		id int8 PRIMARY KEY,
		name varchar(255),
		specie varchar(255)
	)
*/
package ttable2

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/lib/tests/tutils/testtable"
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
	C0IdOrder     tabdef.Order = 0
	C1Name        string       = "name"
	C1NameOrder   tabdef.Order = 1
	C2Specie      string       = "specie"
	C2SpecieOrder tabdef.Order = 2
)

var (
	C0IdType     = hgtype.NewInt8(rawtype.Args{})
	C1NameType   = hgtype.NewStr(rawtype.Args{Length: 255, VarLen: true, UTF8: true})
	C2SpecieType = hgtype.NewStr(rawtype.Args{Length: 255, VarLen: true, UTF8: true})
)
