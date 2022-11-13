package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/analyser/internal/queryerr"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/backend/internal/shared/qctx"
)

var Insert = insert{}

type insert struct{}

func (i insert) Analyse(
	node pnode.InsertNode,
	ctx qctx.QueryCtx,
) (anode.Insert, error) {
	insertNode := anode.Insert{}
	table, err := ctx.QTCtx.GetTableByName(node.Table.TableName)
	if err != nil {
		return insertNode, err
	}
	insertNode.Table, err = ctx.QTCtx.GetOrCreateQTableId(node.Table.TableAlias, table)
	if err != nil {
		return insertNode, err
	}

	if err = i.analyseColumns(node, &insertNode, ctx); err != nil {
		return insertNode, err
	}

	rows, err := InsertRows.Analyse(node.Rows, insertNode.Columns, ctx)
	if err != nil {
		return anode.Insert{}, err
	}

	insertNode.Rows = rows
	return insertNode, nil
}

func (i insert) analyseColumns(
	node pnode.InsertNode,
	insertNode *anode.Insert,
	ctx qctx.QueryCtx,
) error {
	tableDef := ctx.QTCtx.GetTableByQTableId(insertNode.Table)

	if node.ColNames == nil {
		colCount := tableDef.ColumnCount()
		insertNode.Columns = make([]qctx.QColumnId, colCount)
		for j := uint16(0); j < colCount; j++ {
			insertNode.Columns[j] = qctx.QColumnId{
				QTableId: insertNode.Table,
				ColOrder: j,
			}
		}
		return nil
	}

	colNames := node.ColNames
	columns := make([]qctx.QColumnId, len(colNames))
	for j, colName := range colNames {
		colDef, ok := tableDef.ColumnByName(colName)
		if !ok {
			return queryerr.ColumnNotExist(colName, tableDef.Name())
		}
		columns[j] = qctx.QColumnId{QTableId: insertNode.Table, ColOrder: colDef.Order()}
	}
	insertNode.Columns = columns
	return nil
}
