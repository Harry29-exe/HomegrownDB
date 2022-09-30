package ttable1

import "HomegrownDB/dbsystem/schema/column"

const (
	TableName = "awesome_table1"

	C0AwesomeKey      string       = "awesome_key"
	C0AwesomeKeyOrder column.Order = 0

	C1NullableCol      string       = "nullable_col"
	C1NullableColOrder column.Order = 1

	C2NonNullColl      string       = "non_null_coll"
	C2NonNullCollOrder column.Order = 2
)
