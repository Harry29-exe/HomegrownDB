package difinitions

type Column struct {
	Name string
	Id   ColumnId
	Type ColumnType
	// Offset -1 means that offset can not be calculated
	// because one of previous columns was varying size
	// or null
	Offset int32

	Nullable      bool
	Autoincrement bool
}

type ColumnType[T any] interface {
	SerializeData(data T) []byte
	DeserializeData(data []byte) T
}

type ColumnValue interface {
	AsBytes() []byte
	Value() any

	EqualsBytes(value []byte) bool
	Equals(value any) bool

	CompareBytes(value []byte) bool
	Compare(value any) bool
}

type ColumnParser interface {
	Skip(data []byte) []byte
	Parse(data []byte) (ColumnValue, []byte)
}

type ColumnSerializer interface {
}
