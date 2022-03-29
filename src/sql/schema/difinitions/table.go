package difinitions

type Table interface {
	ColumnId(name string) ColumnId
	ColumnsIds(names []string) []ColumnId

	ColumnParsers(ids []ColumnId) []ColumnParser
	ColumnSerializers(ids []ColumnId) []ColumnSerializer
}

type ColumnId = uint16
