package exenode

import (
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/tx"
)

type Builder interface {
	Build(node plan.Node, ctx ExeCtx) ExeNode
}

type ExeCtx = *exeCtx

type exeCtx struct {
	Tx     *tx.Ctx
	Buff   buffer.SharedBuffer
	Stores Stores
}

type Stores struct {
	Fsm   fsm.Store
	Table table.Store
}

func Build(node plan.Node, ctx ExeCtx) ExeNode {
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
