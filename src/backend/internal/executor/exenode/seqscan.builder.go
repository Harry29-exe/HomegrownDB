package exenode

import (
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/table"
)

var SeqScanBuilder Builder = seqScanBuilder{}

type seqScanBuilder struct{}

func (ssb seqScanBuilder) Build(node plan.Node, ctx BuildCtx) ExeNode {
	seqScanNode := node.(plan.SeqScan)
	return NewSeqScan(
		table.DBTableStore.Table(seqScanNode.Table),
		access.DBTableIOStore.TableIO(seqScanNode.Table),
		buffer.DBSharedBuffer,
	)
}
