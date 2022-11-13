package exenode

import (
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/table"
)

var SeqScanBuilder Builder = seqScanBuilder{}

type seqScanBuilder struct{}

func (ssb seqScanBuilder) Build(node plan.Node, ctx ExeCtx) ExeNode {
	seqScanNode := node.(plan.SeqScan)
	return NewSeqScan(
		//todo table should be accessed via plan (executor should lock accordingly all tables specified in plan before execution start)
		ctx.Stores.Table.AccessTable(seqScanNode.Table, table.RLockMode),
		buffer.DBSharedBuffer,
	)
}
