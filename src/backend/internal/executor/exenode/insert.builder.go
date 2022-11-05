package exenode

import (
	"HomegrownDB/backend/internal/planer/plan"
)

func init() {
	exeNodeBuilders[plan.InsertValuesSrcNode] = insertBuilder{}
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
	insertExeNode.rowSrc = Build(insertPlan.Src, ctx)

	return insertExeNode
}
