package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/internal/queryerr"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
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

func (p ColumnTypesPattern) RowMatches(node pnode.InsertingRow, ctx *tx.Ctx) error {
	if len(p.Types) != len(node.Values) {
		return queryerr.NewPatternMatchLenError(len(p.Types), len(node.Values),
			ctx.CurrentQuery.Reconstruct(node.NodeStartTokenIndex, node.NodeEndTokenIndex))
	}
	for i, value := range node.Values {
		if !value.IsAssignableTo(p.Types[i]) {
			return queryerr.NewPatternMatchError(
				value.TypeAsStr(), value.VasStr(), p.Types[i],
				ctx.CurrentQuery.Reconstruct(node.NodeStartTokenIndex, node.NodeEndTokenIndex))
		}
	}

	return nil
}
