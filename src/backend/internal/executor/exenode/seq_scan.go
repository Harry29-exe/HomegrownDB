package exenode

import (
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/access/buffer"
	dbbs2 "HomegrownDB/dbsystem/access/dbbs"
	"HomegrownDB/dbsystem/schema/table"
	page2 "HomegrownDB/dbsystem/storage/page"
)

func NewSeqScan(table table.Definition, tableDataIO access.TableDataIO, buffer buffer.SharedBuffer) *SeqScan {
	return &SeqScan{
		tableDef: table,
		tableIO:  tableDataIO,
		buffer:   buffer,
	}
}

var _ ExeNode = &SeqScan{}

type SeqScan struct {
	tableDef table.Definition
	tableIO  access.TableDataIO
	buffer   buffer.SharedBuffer

	page  page2.Id
	tuple page2.TupleIndex

	hasNext bool
}

func (s *SeqScan) SetSource(source []ExeNode) {
	panic("This node can not have source, it's programmers error")
}

func (s *SeqScan) Free() {
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) HasNext() bool {
	return s.hasNext
}

func (s *SeqScan) Next() dbbs2.QRow {
	tag := buffer.PageTag{PageId: s.page, Relation: s.tableDef.RelationId()}
	rPage, err := s.buffer.RPage(tag)
	if err != nil {
		panic("")
	}
	tablePage := page2.Adapter.TablePage(rPage, s.tableDef)
	defer buffer.DBSharedBuffer.ReleaseRPage(tag)
	tuple := tablePage.Tuple(s.tuple)

	tCount := tablePage.TupleCount()
	if tCount == s.tuple+1 {
		s.tuple = 0
		s.page += 1
		if s.page == s.tableIO.PageCount() {
			s.hasNext = false
		}
	}

	return dbbs2.NewQRowFromTuple(tuple)
}

func (s *SeqScan) NextBatch() []dbbs2.QRow {
	tag := buffer.PageTag{PageId: s.page, Relation: s.tableDef.RelationId()}
	rPage, err := buffer.DBSharedBuffer.RPage(tag)
	if err != nil {
		panic("")
	}

	defer buffer.DBSharedBuffer.ReleaseRPage(tag)
	tCount := rPage.TupleCount()
	rows := make([]dbbs2.QRow, tCount)
	for i := uint16(0); i < tCount; i++ {
		rows[i] = dbbs2.NewQRowFromTuple(rPage.Tuple(i))
	}

	s.page += 1
	if s.page == s.tableIO.PageCount() {
		s.hasNext = false
	}

	return rows
}

func (s *SeqScan) All() []dbbs2.QRow {
	tuplesPerPageEstimate := uint32(page2.Size) / (uint32(s.tableDef.ColumnCount()) * 5)
	rows := make([]dbbs2.QRow, s.tableIO.PageCount()*tuplesPerPageEstimate)
	for s.page < s.tableIO.PageCount() {
		rows = s.readPageWhileReadingAll(rows)

		s.page += 1
	}
	s.hasNext = false

	return rows
}

func (s *SeqScan) readPageWhileReadingAll(rows []dbbs2.QRow) []dbbs2.QRow {
	tag := buffer.PageTag{PageId: s.page, Relation: s.tableDef.RelationId()}
	rPage, err := buffer.DBSharedBuffer.RPage(tag)
	if err != nil {
		panic("")
	}

	defer buffer.DBSharedBuffer.ReleaseRPage(tag)
	tCount := rPage.TupleCount()
	for i := uint16(0); i < tCount; i++ {
		rows = append(rows, dbbs2.NewQRowFromTuple(rPage.Tuple(i)))
	}

	return rows
}

var SeqScanBuilder = seqScanBuilder{}

type seqScanBuilder struct{}

func (ssb seqScanBuilder) Build(node plan.Node) ExeNode {
	seqScanNode := node.(plan.SeqScan)
	return NewSeqScan(table.DBTableStore.Table(seqScanNode.Table), access.DBTableIOStore.TableIO(seqScanNode.Table), buffer.DBSharedBuffer)
}
