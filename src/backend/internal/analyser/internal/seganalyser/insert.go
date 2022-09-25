package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/analyser/internal/queryerr"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/tx"
)

var Insert = insert{}

type insert struct{}

func (i insert) Analyse(
	node pnode.InsertNode,
	ctx *tx.Ctx,
) (anode.Insert, error) {
	insertNode := anode.Insert{}
	table, err := ctx.GetTable(node.Table.TableName)
	if err != nil {
		return insertNode, err
	}
	insertNode.Table = anode.Table{
		Def:      table,
		QTableId: ctx.CurrentQuery.NextQTableId(table.TableId()),
		Alias:    node.Table.TableAlias,
	}

	pattern, err := i.analyseColumns(node, &insertNode)
	if err != nil {
		return insertNode, err
	}

	rows, err := InsertRows.Analyse(node.Rows, pattern, ctx)
	if err != nil {
		return anode.Insert{}, err
	}

	insertNode.Rows = rows
	return insertNode, nil
}

func (i insert) analyseColumns(
	node pnode.InsertNode,
	insertNode *anode.Insert,
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
