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
		AddColumn(C0Id, false, hgtype.TypeInt8, nil).
		AddColumn(C1Age, true, hgtype.TypeInt8, nil).
		AddColumn(C2Name, true, hgtype.TypeStr, nil).
		AddColumn(C3Surname, true, hgtype.TypeStr, nil).
		GetTable()

	return testtable.NewTestTable(table, t)
}

const (
	TableName = "users"

	C0Id      string         = "id"
	C0IdOrder column.Order   = 0
	C0IdType  hgtype.TypeTag = hgtype.TypeInt8

	C1Age      string         = "age"
	C1AgeOrder column.Order   = 1
	C1AgeType  hgtype.TypeTag = hgtype.TypeInt8

	C2Name      string         = "name"
	C2NameOrder column.Order   = 2
	C2NameType  hgtype.TypeTag = hgtype.TypeStr

	C3Surname      string         = "surname"
	C3SurnameOrder column.Order   = 3
	C3SurnameType  hgtype.TypeTag = hgtype.TypeStr
)
