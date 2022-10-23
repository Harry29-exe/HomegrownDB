package page

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/db"
	"HomegrownDB/dbsystem/schema/table"
)

type RPage interface {
	Header() []byte
	Data() []byte
	RelationID() db.RelationID
}

type TableRPage interface {
	Tuple(tupleIndex uint16) Tuple
	TupleCount() uint16
	FreeSpace() uint16

	Data() []byte
}

type TableWPage interface {
	TableRPage

	InsertTuple(data []byte) error
	UpdateTuple(tIndex TupleIndex, tuple []byte)
	DeleteTuple(tIndex TupleIndex)
}

type Id = uint32

const IdSize = 4

type Tag struct {
	PageId   Id
	Relation db.RelationID
}

func NewPageTag(pageIndex Id, tableDef table.Definition) Tag {
	return Tag{
		PageId:   pageIndex,
		Relation: tableDef.RelationId(),
	}
}

const Size uint16 = dbsystem.PageSize
