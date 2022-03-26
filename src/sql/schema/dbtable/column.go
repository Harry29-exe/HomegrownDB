package dbtable

import "HomegrownDB/sql/schema"

type Column struct {
	Name string
	Id   schema.ColumnId
	Type ColumnType
	// Offset -1 means that offset can not be calculated
	// because one of previous columns was varying size
	// or null
	Offset int32

	Nullable      bool
	Autoincrement bool
}
