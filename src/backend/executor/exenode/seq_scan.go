package exenode

import (
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/io/buffer"
	"HomegrownDB/dbsystem/schema/table"
)

type SeqScan struct {
	table  table.Definition
	buffer buffer.DBSharedBuffer

	currentData []bdata.Tuple
	nextPage    bdata.PageId
}

func NewSeqScan(table table.Definition, buffer buffer.DBSharedBuffer) *SeqScan {
	return &SeqScan{
		table:  table,
		buffer: buffer,
	}
}

func (s *SeqScan) HasNext() bool {
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) Next() bdata.Tuple {
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) NextBatch() []bdata.Tuple {
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) All() []bdata.Tuple {
	//TODO implement me
	panic("implement me")
}

func (s *SeqScan) loadDataFromNextPage() {
	//todo implement me
	panic("Not implemented")
}
