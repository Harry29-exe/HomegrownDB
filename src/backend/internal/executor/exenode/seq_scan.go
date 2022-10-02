package exenode

import (
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/schema/table"
)

func NewSeqScan(table table.Definition, tableDataIO access.TableDataIO, buffer buffer.DBSharedBuffer) *SeqScan {
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
	buffer   buffer.DBSharedBuffer

	page  dbbs.PageId
	tuple dbbs.TupleIndex

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

func (s *SeqScan) Next() dbbs.QRow {
	tag := dbbs.PageTag{PageId: s.page, TableId: s.tableDef.TableId()}
	rPage, err := s.buffer.RPage(tag)
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

	return dbbs.NewQRowFromTuple(tuple)
}

func (s *SeqScan) NextBatch() []dbbs.QRow {
	tag := dbbs.PageTag{PageId: s.page, TableId: s.tableDef.TableId()}
	rPage, err := buffer.SharedBuffer.RPage(tag)
	if err != nil {
		panic("")
	}

	defer buffer.SharedBuffer.ReleaseRPage(tag)
	tCount := rPage.TupleCount()
	rows := make([]dbbs.QRow, tCount)
	for i := uint16(0); i < tCount; i++ {
		rows[i] = dbbs.NewQRowFromTuple(rPage.Tuple(i))
	}

	s.page += 1
	if s.page == s.tableIO.PageCount() {
		s.hasNext = false
	}

	return rows
}

func (s *SeqScan) All() []dbbs.QRow {
	tuplesPerPageEstimate := uint32(dbbs.PageSize) / (uint32(s.tableDef.ColumnCount()) * 5)
	rows := make([]dbbs.QRow, s.tableIO.PageCount()*tuplesPerPageEstimate)
	for s.page < s.tableIO.PageCount() {
		rows = s.readPageWhileReadingAll(rows)

		s.page += 1
	}
	s.hasNext = false

	return rows
}

func (s *SeqScan) readPageWhileReadingAll(rows []dbbs.QRow) []dbbs.QRow {
	tag := dbbs.PageTag{PageId: s.page, TableId: s.tableDef.TableId()}
	rPage, err := buffer.SharedBuffer.RPage(tag)
	if err != nil {
		panic("")
	}

	defer buffer.SharedBuffer.ReleaseRPage(tag)
	tCount := rPage.TupleCount()
	for i := uint16(0); i < tCount; i++ {
		rows = append(rows, dbbs.NewQRowFromTuple(rPage.Tuple(i)))
	}

	return rows
}

var SeqScanBuilder = seqScanBuilder{}

type seqScanBuilder struct{}

func (ssb seqScanBuilder) Build(node plan.Node) ExeNode {
	seqScanNode := node.(plan.SeqScan)
	return NewSeqScan(table.DBTableStore.Table(seqScanNode.Table), access.DBTableIOStore.TableIO(seqScanNode.Table), buffer.SharedBuffer)
}
