package dbbs

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/db"
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
	PageId   PageId
	Relation db.RelationID
}

func NewPageTag(pageIndex PageId, tableDef table.Definition) PageTag {
	return PageTag{
		PageId:   pageIndex,
		Relation: tableDef.RelationId(),
	}
}

const PageSize uint16 = dbsystem.PageSize
