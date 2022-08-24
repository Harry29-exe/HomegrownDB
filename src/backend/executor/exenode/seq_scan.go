package exenode

import (
	"HomegrownDB/backend/executor/exenode/internal/data"
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/io/buffer"
	"HomegrownDB/dbsystem/schema/table"
)

type SeqScan struct {
	table  table.Definition
	buffer buffer.DBSharedBuffer
	holder data.RowBuffer

	page  bdata.PageId
	tuple bdata.TupleIndex

	pageCount bdata.PageId
	hasNext   bool
}

func NewSeqScan(table table.Definition, pageCount uint32, buffer buffer.DBSharedBuffer) *SeqScan {
	return &SeqScan{
		table:     table,
		pageCount: pageCount,
		buffer:    buffer,
	}
}

func (s *SeqScan) Init(options InitOptions) data.RowBuffer {
	s.holder = data.NewBaseRowHolder(data.GlobalSlotBuffer, []table.Definition{s.table})

	return s.holder
}

func (s *SeqScan) Free() {
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) HasNext() bool {
	return s.hasNext
}

func (s *SeqScan) Next() data.Row {
	tag := bdata.PageTag{PageId: s.page, TableId: s.table.TableId()}
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
		if s.page == s.pageCount {
			s.hasNext = false
		}
	}

	return data.NewRow([]bdata.Tuple{tuple}, s.holder)
}

func (s *SeqScan) NextBatch() []data.Row {
	tag := bdata.PageTag{PageId: s.page, TableId: s.table.TableId()}
	rPage, err := buffer.SharedBuffer.RPage(tag)
	if err != nil {
		panic("")
	}

	defer buffer.SharedBuffer.ReleaseRPage(tag)
	tCount := rPage.TupleCount()
	rows := make([]data.Row, tCount)
	for i := uint16(0); i < tCount; i++ {
		rows[i] = data.NewRow([]bdata.Tuple{rPage.Tuple(i)}, s.holder)
	}

	s.page += 1
	if s.page == s.pageCount {
		s.hasNext = false
	}

	return rows
}

func (s *SeqScan) All() []data.Row {
	//TODO implement me
	panic("implement me")
}
