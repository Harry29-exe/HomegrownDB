package exenode

import (
	"HomegrownDB/backend/internal/shared/query"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/tpage"
)

func NewSeqScan(table table.RDefinition, buffer buffer.SharedBuffer) *SeqScan {
	return &SeqScan{
		tableDef: table,
		buffer:   buffer,
	}
}

var _ ExeNode = &SeqScan{}

type SeqScan struct {
	tableDef table.RDefinition
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

func (s *SeqScan) Next() query.QRow {
	//todo implement me
	panic("Not implemented")
	//rPage, err := s.buffer.RTablePage(s.tableDef, s.page)
	//if err != nil {
	//	panic("")
	//}
	//
	//defer buffer.DBSharedBuffer.RPageRelease(rPage.PageTag())
	//tuple := rPage.Tuple(s.tuple)
	//
	//tCount := rPage.TupleCount()
	//if tCount == s.tuple+1 {
	//	s.tuple = 0
	//	s.page += 1
	//	if s.page == s.tableIO.PageCount() {
	//		s.hasNext = false
	//	}
	//}
	//
	//return anlsr.NewQRowFromTuple(tuple)
}

func (s *SeqScan) NextBatch() []query.QRow {
	//todo implement me
	panic("Not implemented")
	//rPage, err := buffer.DBSharedBuffer.RTablePage(s.tableDef, s.page)
	//if err != nil {
	//	panic("")
	//}
	//
	//defer buffer.DBSharedBuffer.RPageRelease(rPage.PageTag())
	//tCount := rPage.TupleCount()
	//rows := make([]anlsr.QRow, tCount)
	//for i := uint16(0); i < tCount; i++ {
	//	rows[i] = anlsr.NewQRowFromTuple(rPage.Tuple(i))
	//}
	//
	//s.page += 1
	//if s.page == s.tableIO.PageCount() {
	//	s.hasNext = false
	//}
	//
	//return rows
}

func (s *SeqScan) All() []query.QRow {
	//todo implement me
	panic("Not implemented")
	//tuplesPerPageEstimate := uint32(page.Size) / (uint32(s.tableDef.ColumnCount()) * 5)
	//rows := make([]anlsr.QRow, s.tableIO.PageCount()*tuplesPerPageEstimate)
	//for s.page < s.tableIO.PageCount() {
	//	rows = s.readPageWhileReadingAll(rows)
	//
	//	s.page += 1
	//}
	//s.hasNext = false
	//
	//return rows
}

func (s *SeqScan) readPageWhileReadingAll(rows []query.QRow) []query.QRow {
	rPage, err := buffer.DBSharedBuffer.RTablePage(s.tableDef, s.page)
	if err != nil {
		panic("")
	}

	defer buffer.DBSharedBuffer.RPageRelease(rPage.PageTag())
	tCount := rPage.TupleCount()
	for i := uint16(0); i < tCount; i++ {
		rows = append(rows, query.NewQRowFromTuple(rPage.Tuple(i)))
	}

	return rows
}
