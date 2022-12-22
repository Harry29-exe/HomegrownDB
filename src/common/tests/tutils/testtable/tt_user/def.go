/*
Package tt_user is following table definition

	CREATE TABLE users (
		id int8 PRIMARY KEY,
		age int8,
		name varchar(255),
		surname varchar(255),
	)
*/
package tt_user

import (
	"HomegrownDB/common/tests/tutils/testtable"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/schema/column"
	"testing"
)

func Def(t *testing.T) testtable.TestTable {
	table := testtable.NewTestTableBuilder(TableName).
		AddColumn(C0Id, false, hgtype.TypeInt8, hgtype.Args{}).
		AddColumn(C1Age, true, hgtype.TypeInt8, hgtype.Args{}).
		AddColumn(C2Name, true, hgtype.TypeStr, hgtype.Args{}).
		AddColumn(C3Surname, true, hgtype.TypeStr, hgtype.Args{}).
		GetTable()

	return testtable.NewTestTable(table, t)
}

const (
	TableName = "users"

	C0Id      string       = "id"
	C0IdOrder column.Order = 0

	C1Age      string       = "age"
	C1AgeOrder column.Order = 1

	C2Name      string       = "name"
	C2NameOrder column.Order = 2

	C3Surname      string       = "surname"
	C3SurnameOrder column.Order = 3
)

var (
	C0IdType      = hgtype.NewInt8(hgtype.Args{})
	C1AgeType     = hgtype.NewInt8(hgtype.Args{})
	C2NameType    = hgtype.NewStr(hgtype.Args{Length: 255, VarLen: true, UTF8: true})
	C3SurnameType = hgtype.NewStr(hgtype.Args{Length: 255, VarLen: true, UTF8: true})
)
