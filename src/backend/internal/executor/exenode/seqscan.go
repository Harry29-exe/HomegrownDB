package exenode

import (
	dbbs2 "HomegrownDB/backend/internal/shared/query"
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/table"
	page "HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/tpage"
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

	page  page.Id
	tuple tpage.TupleIndex

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
	rPage, err := s.buffer.RTablePage(s.tableDef, s.page)
	if err != nil {
		panic("")
	}

	defer buffer.DBSharedBuffer.RPageRelease(rPage.PageTag())
	tuple := rPage.Tuple(s.tuple)

	tCount := rPage.TupleCount()
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
	rPage, err := buffer.DBSharedBuffer.RTablePage(s.tableDef, s.page)
	if err != nil {
		panic("")
	}

	defer buffer.DBSharedBuffer.RPageRelease(rPage.PageTag())
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
	tuplesPerPageEstimate := uint32(page.Size) / (uint32(s.tableDef.ColumnCount()) * 5)
	rows := make([]dbbs2.QRow, s.tableIO.PageCount()*tuplesPerPageEstimate)
	for s.page < s.tableIO.PageCount() {
		rows = s.readPageWhileReadingAll(rows)

		s.page += 1
	}
	s.hasNext = false

	return rows
}

func (s *SeqScan) readPageWhileReadingAll(rows []dbbs2.QRow) []dbbs2.QRow {
	rPage, err := buffer.DBSharedBuffer.RTablePage(s.tableDef, s.page)
	if err != nil {
		panic("")
	}

	defer buffer.DBSharedBuffer.RPageRelease(rPage.PageTag())
	tCount := rPage.TupleCount()
	for i := uint16(0); i < tCount; i++ {
		rows = append(rows, dbbs2.NewQRowFromTuple(rPage.Tuple(i)))
	}

	return rows
}
