package seganalyser

import (
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/backend/internal/shared/qctx"
	"HomegrownDB/dbsystem/schema/table"
)

var Tables = tables{}

type tables struct{}

func (t tables) Analise(node []pnode.TableNode, ctx qctx.QueryCtx) ([]qctx.QTableId, error) {
	tablesCount := len(node)
	qTableIds := make([]qctx.QTableId, tablesCount)

	var tableDef table.RDefinition
	var qid qctx.QTableId
	var err error
	for i, tableNode := range node {

		if tableDef, err = ctx.QTCtx.GetTableByName(tableNode.TableName); err != nil {
			return qTableIds, err
		}
		if qid, err = ctx.QTCtx.GetOrCreateQTableId(tableNode.TableAlias, tableDef); err != nil {
			return qTableIds, err
		}

		qTableIds[i] = qid
	}

	return qTableIds, nil
}
