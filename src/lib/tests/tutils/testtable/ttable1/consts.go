package ttable1

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
)

const (
	TableName = "awesome_table1"

	C0AwesomeKey      string       = "awesome_key"
	C0AwesomeKeyOrder reldef.Order = 0

	C1NullableCol      string       = "nullable_col"
	C1NullableColOrder reldef.Order = 1

	C2NonNullColl      string       = "non_null_coll"
	C2NonNullCollOrder reldef.Order = 2
)

var (
	C0AwesomeKeyType  = hgtype.NewInt8(rawtype.Args{})
	C1NullableColType = hgtype.NewInt8(rawtype.Args{})
	C2NonNullCollType = hgtype.NewInt8(rawtype.Args{})
)
