package exenode

import (
	"HomegrownDB/backend/internal/planer/plan"
)

func init() {
	exeNodeBuilders[plan.InsertNode] = insertBuilder{}
}

type insertBuilder struct{}

func (i insertBuilder) Build(node plan.Node, ctx ExeCtx) ExeNode {
	//insertPlan, ok := node.(*plan.Insert)
	//if !ok {
	//	panic("illegal type (expected *plan.Insert)")
	//}

	//insertExeNode := &Insert{
	//	table:  insertPlan.Table,
	//	rowSrc: nil,
	//	txCtx:  ctx.Tx,
	//}
	//columns := make([]column.Id, len(insertPlan.Columns))
	//for j, col := range insertPlan.Columns {
	//	columns[j] = col.Id()
	//}
	//insertExeNode.columns = columns
	//insertExeNode.rowSrc = Build(insertPlan.Src, ctx)

	//return insertExeNode
	//todo implement me
	panic("Not implemented")
}
