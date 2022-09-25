package exenode

import (
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/backend/qrow"
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/schema/table"
)

func NewSeqScan(table table.Definition, tableDataIO access.TableDataIO, buffer buffer.DBSharedBuffer) *SeqScan {
	return &SeqScan{
		tableDef: table,
		tableIO:  tableDataIO,
		buffer:   buffer,
	}
}

type SeqScan struct {
	tableDef table.Definition
	tableIO  access.TableDataIO
	buffer   buffer.DBSharedBuffer
	holder   qrow.RowBuffer

	page  bdata.PageId
	tuple bdata.TupleIndex

	hasNext bool
}

func (s *SeqScan) SetSource(source []ExeNode) {
	panic("This node can not have source, it's programmers error")
}

func (s *SeqScan) Init(options InitOptions) qrow.RowBuffer {
	s.holder = qrow.NewBaseRowHolder(qrow.GlobalSlotBuffer, []table.Definition{s.tableDef})

	return s.holder
}

func (s *SeqScan) Free() {
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) HasNext() bool {
	return s.hasNext
}

func (s *SeqScan) Next() qrow.Row {
	tag := bdata.PageTag{PageId: s.page, TableId: s.tableDef.TableId()}
	rPage, err := buffer.SharedBuffer.RPage(tag)
	if err != nil {
		panic("")
	}
	defer buffer.SharedBuffer.ReleaseRPage(tag)
	tuple := rPage.Tuple(s.tuple)

	tCount := rPage.TupleCount()
	if tCount == s.tuple+1 {
		s.tuple = 0
		s.page += 1
		if s.page == s.tableIO.PageCount() {
			s.hasNext = false
		}
	}

	return qrow.NewRow([]bdata.Tuple{tuple}, s.holder)
}

func (s *SeqScan) NextBatch() []qrow.Row {
	tag := bdata.PageTag{PageId: s.page, TableId: s.tableDef.TableId()}
	rPage, err := buffer.SharedBuffer.RPage(tag)
	if err != nil {
		panic("")
	}

	defer buffer.SharedBuffer.ReleaseRPage(tag)
	tCount := rPage.TupleCount()
	rows := make([]qrow.Row, tCount)
	for i := uint16(0); i < tCount; i++ {
		rows[i] = qrow.NewRow([]bdata.Tuple{rPage.Tuple(i)}, s.holder)
	}

	s.page += 1
	if s.page == s.tableIO.PageCount() {
		s.hasNext = false
	}

	return rows
}

func (s *SeqScan) All() []qrow.Row {
	tuplesPerPageEstimate := uint32(bdata.PageSize) / (uint32(s.tableDef.ColumnCount()) * 5)
	rows := make([]qrow.Row, s.tableIO.PageCount()*tuplesPerPageEstimate)
	for s.page < s.tableIO.PageCount() {
		rows = s.readPageWhileReadingAll(rows)

		s.page += 1
	}
	s.hasNext = false

	return rows
}

func (s *SeqScan) readPageWhileReadingAll(rows []qrow.Row) []qrow.Row {
	tag := bdata.PageTag{PageId: s.page, TableId: s.tableDef.TableId()}
	rPage, err := buffer.SharedBuffer.RPage(tag)
	if err != nil {
		panic("")
	}

	defer buffer.SharedBuffer.ReleaseRPage(tag)
	tCount := rPage.TupleCount()
	for i := uint16(0); i < tCount; i++ {
		rows = append(rows, qrow.NewRow([]bdata.Tuple{rPage.Tuple(i)}, s.holder))
	}

	return rows
}

var SeqScanBuilder = seqScanBuilder{}

type seqScanBuilder struct{}

func (ssb seqScanBuilder) Build(node plan.Node) ExeNode {
	seqScanNode := node.(plan.SeqScan)
	return NewSeqScan(table.DBTableStore.Table(seqScanNode.Table), access.DBTableIOStore.TableIO(seqScanNode.Table), buffer.SharedBuffer)
}
