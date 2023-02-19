package ttable1

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef/tabdef"
)

const (
	TableName = "awesome_table1"

	C0AwesomeKey      string       = "awesome_key"
	C0AwesomeKeyOrder tabdef.Order = 0

	C1NullableCol      string       = "nullable_col"
	C1NullableColOrder tabdef.Order = 1

	C2NonNullColl      string       = "non_null_coll"
	C2NonNullCollOrder tabdef.Order = 2
)

var (
	C0AwesomeKeyType  = hgtype.NewInt8(rawtype.Args{})
	C1NullableColType = hgtype.NewInt8(rawtype.Args{})
	C2NonNullCollType = hgtype.NewInt8(rawtype.Args{})
)
