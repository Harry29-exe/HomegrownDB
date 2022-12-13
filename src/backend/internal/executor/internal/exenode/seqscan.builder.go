package exenode

import (
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/dbsystem/access/buffer"
)

var SeqScanBuilder Builder = seqScanBuilder{}

type seqScanBuilder struct{}

func (ssb seqScanBuilder) Build(node plan.Node, ctx ExeCtx) ExeNode {
	//seqScanNode := node.(plan.SeqScan)
	return NewSeqScan(
		nil,
		buffer.DBSharedBuffer,
	)
}
