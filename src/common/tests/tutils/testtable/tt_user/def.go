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
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
	"testing"
)

func Def(t *testing.T) testtable.TestTable {
	table := testtable.NewTestTableBuilder(TableName).
		AddColumn(C0Id, false, ctype.TypeInt8, nil).
		AddColumn(C1Age, true, ctype.TypeInt8, nil).
		AddColumn(C2Name, true, ctype.TypeStr, nil).
		AddColumn(C3Surname, true, ctype.TypeStr, nil).
		GetTable()

	return testtable.NewTestTable(table, t)
}

const (
	TableName = "users"

	C0Id      string       = "id"
	C0IdOrder column.Order = 0
	C0IdType  ctype.Type   = ctype.TypeInt8

	C1Age      string       = "age"
	C1AgeOrder column.Order = 1
	C1AgeType  ctype.Type   = ctype.TypeInt8

	C2Name      string       = "name"
	C2NameOrder column.Order = 2
	C2NameType  ctype.Type   = ctype.TypeStr

	C3Surname      string       = "surname"
	C3SurnameOrder column.Order = 3
	C3SurnameType  ctype.Type   = ctype.TypeStr
)
