package ttable1

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/relation/table/column"
)

const (
	TableName = "awesome_table1"

	C0AwesomeKey      string       = "awesome_key"
	C0AwesomeKeyOrder column.Order = 0

	C1NullableCol      string       = "nullable_col"
	C1NullableColOrder column.Order = 1

	C2NonNullColl      string       = "non_null_coll"
	C2NonNullCollOrder column.Order = 2
)

var (
	C0AwesomeKeyType  = hgtype.NewInt8(hgtype.Args{})
	C1NullableColType = hgtype.NewInt8(hgtype.Args{})
	C2NonNullCollType = hgtype.NewInt8(hgtype.Args{})
)
