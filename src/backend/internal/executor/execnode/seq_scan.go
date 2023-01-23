package execnode

import (
	"HomegrownDB/backend/internal/executor/exinfr"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

var _ Builder = seqScanBuilder{}

type seqScanBuilder struct{}

func (s seqScanBuilder) Create(plan node.Plan, ctx exinfr.ExCtx) ExecNode {
	seqScanPlan := plan.(node.SeqScan)
	scanTable := ctx.GetRTE(seqScanPlan.RteId).Ref
	return &SeqScan{
		scan: scan{
			Plan: plan,
			Tx:   ctx.Tx,
		},
		Plan:          seqScanPlan,
		OutputPattern: page.PatternFromTable(scanTable),
		buff:          ctx.Buff,
		table:         scanTable,
		nextPageId:    0,
		nextTupleId:   0,
		done:          false,
	}
}

var _ ExecNode = &SeqScan{}

type SeqScan struct {
	scan
	Plan          node.SeqScan
	OutputPattern page.TuplePattern

	txCtx tx.Tx
	buff  buffer.SharedBuffer
	table table.RDefinition

	nextPageId  page.Id
	nextTupleId page.TupleIndex
	done        bool
}

//todo !!!!Next and HasNext are poc, they are awful for performance!!!!

func (s *SeqScan) Next() page.Tuple {
	rPage, err := s.buff.RTablePage(s.table, s.nextPageId)
	if err != nil {
		s.done = true
	}
	defer s.buff.RPageRelease(rPage.PageTag())
	tuple := rPage.Tuple(s.nextTupleId)
	outputTuple := s.createOutputTuple(tuple)
	s.nextTupleId++
	if s.nextTupleId == rPage.TupleCount() {
		s.nextTupleId = 0
		s.nextPageId++
	}

	return outputTuple
}

func (s *SeqScan) HasNext() bool {
	rPage, err := s.buff.RTablePage(s.table, s.nextPageId)
	if err != nil {
		return false
	}
	defer s.buff.RPageRelease(rPage.PageTag())
	return rPage.TupleCount() > s.nextTupleId
}

func (s *SeqScan) Init(plan node.Plan) error {
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) Close() error {
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) mapTuple(tuple page.WTuple) page.WTuple {
	//todo implement me
	panic("Not implemented")
}
