package execnode

import (
	"HomegrownDB/backend/internal/executor/exinfr"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dpage"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/tpage"
	"HomegrownDB/dbsystem/tx"
)

var _ Builder = seqScanBuilder{}

type seqScanBuilder struct{}

func (s seqScanBuilder) Create(plan node.Plan, ctx exinfr.ExCtx) ExecNode {
	seqScanPlan := plan.(node.SeqScan)
	scanTable := ctx.GetRTE(seqScanPlan.RteId).Ref
	return &SeqScan{
		Plan:          seqScanPlan,
		OutputPattern: dpage.NewPatternFromTable(scanTable),
		txCtx:         ctx.Tx,
		buff:          ctx.Buff,
		table:         scanTable,
		nextPageId:    0,
		nextTupleId:   0,
		done:          false,
	}
}

var _ ExecNode = &SeqScan{}

type SeqScan struct {
	Plan          node.SeqScan
	OutputPattern *dpage.TuplePattern

	txCtx tx.Tx
	buff  buffer.SharedBuffer
	table table.RDefinition

	nextPageId  page.Id
	nextTupleId tpage.TupleIndex
	done        bool

}

func (s SeqScan) Next() dpage.Tuple {
	rPage, err := s.buff.RTablePage(s.table, s.nextPageId)
	if err != nil {
		s.done = true
	}

}

func (s SeqScan) HasNext() bool {
	//TODO implement me
	panic("implement me")
}

func (s SeqScan) Init(plan node.Plan) error {
	//TODO implement me
	panic("implement me")
}
