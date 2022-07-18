package bdata

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/schema/table"
)

type RPage interface {
	Tuple(tupleIndex uint16) Tuple
	TupleCount() uint16
	FreeSpace() uint16
}

type WPage interface {
	RPage

	InsertTuple(data []byte) error
	UpdateTuple(tIndex TupleIndex, tuple []byte)
	DeleteTuple(tIndex TupleIndex)
	Data() []byte
}

type PageId = uint32

const PageIdSize = 4

type PageTag struct {
	PageId  PageId
	TableId table.Id
}

func NewPageTag(pageIndex PageId, def table.Definition) PageTag {
	return PageTag{
		PageId:  pageIndex,
		TableId: def.TableId(),
	}
}

const PageSize uint16 = dbsystem.PageSize
