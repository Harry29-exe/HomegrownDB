package exenode

import (
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/dbsystem/schema/column"
)

func init() {
	exeNodeBuilders[plan.InsertNode] = insertBuilder{}
}

type insertBuilder struct{}

func (i insertBuilder) Build(node plan.Node, ctx BuildCtx) ExeNode {
	insertPlan, ok := node.(*plan.Insert)
	if !ok {
		panic("illegal type (expected *plan.Insert)")
	}

	insertExeNode := &Insert{
		table:  insertPlan.Table,
		rowSrc: nil,
		txCtx:  ctx.tx,
	}
	columns := make([]column.Id, len(insertPlan.Columns))
	for j, col := range insertPlan.Columns {
		columns[j] = col.Id()
	}
	insertExeNode.columns = columns
	insertExeNode.rowSrc = Build(insertPlan.Src, ctx)

	return insertExeNode
}
