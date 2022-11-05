package exenode

import (
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
)

type Builder interface {
	Build(node plan.Node, ctx BuildCtx) ExeNode
}

type BuildCtx = *buildCtx

type buildCtx struct {
	tx         *tx.Ctx
	buff       buffer.SharedBuffer
	tableStore table.Store
}

func Build(node plan.Node, ctx BuildCtx) ExeNode {
	builder, ok := exeNodeBuilders[node.Type()]
	if !ok {
		//todo implement me
		panic("Not implemented")
	}

	return builder.Build(node, ctx)
}

var exeNodeBuilders = map[plan.NodeType]Builder{
	plan.SeqScanNode: SeqScanBuilder,
}
