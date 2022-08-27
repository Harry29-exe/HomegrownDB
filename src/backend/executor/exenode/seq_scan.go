package exenode

import (
	qrow2 "HomegrownDB/backend/executor/qrow"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/io"
	"HomegrownDB/dbsystem/io/buffer"
	"HomegrownDB/dbsystem/schema/table"
)

type SeqScan struct {
	tableDef table.Definition
	tableIO  io.TableDataIO
	buffer   buffer.DBSharedBuffer
	holder   qrow2.RowBuffer

	page  bdata.PageId
	tuple bdata.TupleIndex

	hasNext bool
}

func NewSeqScan(table table.Definition, tableDataIO io.TableDataIO, buffer buffer.DBSharedBuffer) *SeqScan {
	return &SeqScan{
		tableDef: table,
		tableIO:  tableDataIO,
		buffer:   buffer,
	}
}

func (s *SeqScan) Init(options InitOptions) qrow2.RowBuffer {
	s.holder = qrow2.NewBaseRowHolder(qrow2.GlobalSlotBuffer, []table.Definition{s.tableDef})

	return s.holder
}

func (s *SeqScan) Free() {
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) HasNext() bool {
	return s.hasNext
}

func (s *SeqScan) Next() qrow2.Row {
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

	return qrow2.NewRow([]bdata.Tuple{tuple}, s.holder)
}

func (s *SeqScan) NextBatch() []qrow2.Row {
	tag := bdata.PageTag{PageId: s.page, TableId: s.tableDef.TableId()}
	rPage, err := buffer.SharedBuffer.RPage(tag)
	if err != nil {
		panic("")
	}

	defer buffer.SharedBuffer.ReleaseRPage(tag)
	tCount := rPage.TupleCount()
	rows := make([]qrow2.Row, tCount)
	for i := uint16(0); i < tCount; i++ {
		rows[i] = qrow2.NewRow([]bdata.Tuple{rPage.Tuple(i)}, s.holder)
	}

	s.page += 1
	if s.page == s.tableIO.PageCount() {
		s.hasNext = false
	}

	return rows
}

func (s *SeqScan) All() []qrow2.Row {
	tuplesPerPageEstimate := uint32(bdata.PageSize) / (uint32(s.tableDef.ColumnCount()) * 5)
	rows := make([]qrow2.Row, s.tableIO.PageCount()*tuplesPerPageEstimate)
	for s.page < s.tableIO.PageCount() {
		rows = s.readPageWhileReadingAll(rows)

		s.page += 1
	}
	s.hasNext = false

	return rows
}

func (s *SeqScan) readPageWhileReadingAll(rows []qrow2.Row) []qrow2.Row {
	tag := bdata.PageTag{PageId: s.page, TableId: s.tableDef.TableId()}
	rPage, err := buffer.SharedBuffer.RPage(tag)
	if err != nil {
		panic("")
	}

	defer buffer.SharedBuffer.ReleaseRPage(tag)
	tCount := rPage.TupleCount()
	for i := uint16(0); i < tCount; i++ {
		rows = append(rows, qrow2.NewRow([]bdata.Tuple{rPage.Tuple(i)}, s.holder))
	}

	return rows
}
