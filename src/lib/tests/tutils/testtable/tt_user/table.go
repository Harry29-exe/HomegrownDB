/*
Package tt_user is following tabdef definition

	CREATE TABLE users (
		id int8 PRIMARY KEY,
		age int8,
		name varchar(255),
		surname varchar(255),
	)
*/
package tt_user

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/lib/tests/tutils/testtable"
	"testing"
)

func Def(t *testing.T) testtable.TestTable {
	table := testtable.NewTestTableBuilder(TableName).
		AddColumn(C0Id, false, C0IdType).
		AddColumn(C1Age, true, C1AgeType).
		AddColumn(C2Name, true, C2NameType).
		AddColumn(C3Surname, true, C3SurnameType).
		GetTable()

	return testtable.NewTestTable(table, t)
}

const (
	TableName = "users"

	C0Id      string       = "id"
	C0IdOrder reldef.Order = 0

	C1Age      string       = "age"
	C1AgeOrder reldef.Order = 1

	C2Name      string       = "name"
	C2NameOrder reldef.Order = 2

	C3Surname      string       = "surname"
	C3SurnameOrder reldef.Order = 3
)

var (
	C0IdType      = hgtype.NewInt8(rawtype.Args{})
	C1AgeType     = hgtype.NewInt8(rawtype.Args{})
	C2NameType    = hgtype.NewStr(rawtype.Args{Length: 255, VarLen: true, UTF8: true})
	C3SurnameType = hgtype.NewStr(rawtype.Args{Length: 255, VarLen: true, UTF8: true})
)
