package seganalyser

import (
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
)

func NewRowCTypesPattern(colsIds []column.OrderId, table table.Definition) ColumnTypesPattern {
	pattern := ColumnTypesPattern{Types: make([]column.Type, len(colsIds))}
	for i, id := range colsIds {
		pattern.Types[i] = table.GetColumn(id).Type()
	}

	return pattern
}

type ColumnTypesPattern struct {
	Types []column.Type
}

func (p ColumnTypesPattern) RowMatches(node pnode.InsertingRow) error {
	if len(p.Types) != len(node.Values) {
		return false
	}
	for i, value := range node.Values {
		if !value.IsAssignableTo(p.Types[i]) {
			return false
		}
	}

	return true
}
