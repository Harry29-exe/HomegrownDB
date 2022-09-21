package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/analyser/internal"
	"HomegrownDB/backend/internal/analyser/internal/queryerr"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/schema/column"
)

var Insert = insert{}

type insert struct{}

func (i insert) Analyse(
	node pnode.InsertNode,
	ctx *internal.AnalyserCtx,
) (anode.Insert, error) {
	insertNode := anode.Insert{}
	table, err := ctx.GetTable(node.Table.TableName)
	if err != nil {
		return insertNode, err
	}
	insertNode.Table = anode.Table{
		Def:      table,
		QTableId: ctx.NextQTableId(table.TableId()),
		Alias:    node.Table.TableAlias,
	}

	_, err = i.analyseColumns(node, insertNode)
	if err != nil {
		return insertNode, err
	}

	//todo implement me
	panic("Not implemented")
}

func (i insert) analyseColumns(
	node pnode.InsertNode,
	insertNode anode.Insert,
) (ColumnTypesPattern, error) {
	tableDef := insertNode.Table.Def

	if node.ColNames == nil {
		colCount := tableDef.ColumnCount()
		insertNode.Columns = make([]column.OrderId, colCount)
		for j := uint16(0); j < colCount; j++ {
			insertNode.Columns[j] = j
		}
		return NewRowCTypesPattern(insertNode.Columns, tableDef), nil
	}

	colNames := node.ColNames
	columns, ok := make([]column.OrderId, len(colNames)), false
	insertNode.Columns = columns
	for j, colName := range colNames {
		columns[j], ok = tableDef.ColumnId(colName)
		if !ok {
			return ColumnTypesPattern{}, queryerr.ColumnNotExist(colName, tableDef.Name())
		}
	}
	return NewRowCTypesPattern(insertNode.Columns, tableDef), nil
}
