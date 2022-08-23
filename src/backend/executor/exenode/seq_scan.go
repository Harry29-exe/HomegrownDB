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

	nextPage bdata.PageId
	holder   data.RowBuffer
}

func NewSeqScan(table table.Definition, buffer buffer.DBSharedBuffer) *SeqScan {
	return &SeqScan{
		table:  table,
		buffer: buffer,
	}
}

func (s *SeqScan) Init(dataHolder data.RowBuffer) {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) NextBatch() []data.Row {
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) All() []data.Row {
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) loadDataFromNextPage() {
	//todo implement me
	panic("Not implemented")
}
