package pnode

type InsertNode struct {
	Table   TableNode
	Columns InsertingColumns
	Values  []InsertingValues
}

type InsertingColumns struct {
	ColumnNames []string
}

type InsertingValues struct {
	Values []any
}
