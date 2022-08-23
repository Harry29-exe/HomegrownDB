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

	page   bdata.PageId
	tuple  bdata.TupleIndex
	holder data.RowBuffer
}

func NewSeqScan(table table.Definition, buffer buffer.DBSharedBuffer) *SeqScan {
	return &SeqScan{
		table:  table,
		buffer: buffer,
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
	//TODO implement me
	panic("implement me")
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
		s.page += 1 //todo check if table has next page if not set some flag 'hasNext' to false
	}

	return data.NewRow([]bdata.Tuple{tuple}, s.holder)
}

func (s *SeqScan) NextBatch() []data.Row {
	rPage, err := buffer.SharedBuffer.RPage(bdata.PageTag{PageId: s.page, TableId: s.table.TableId()})
	if err != nil {
		panic("")
	}
	rPage
}

func (s *SeqScan) All() []data.Row {
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) loadDataFromNextPage() {
	//todo implement me
	panic("Not implemented")
}
