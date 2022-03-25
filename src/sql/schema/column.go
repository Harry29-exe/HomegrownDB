package schema

type Column struct {
	Name string
	Type ColumnType
	// Offset -1 means that offset can not be calculated
	// because one of previous columns was varying size
	Offset int32

	Nullable      bool
	Autoincrement bool
}
