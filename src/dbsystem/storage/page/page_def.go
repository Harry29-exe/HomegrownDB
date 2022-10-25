package page

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/schema/table"
)

type Page interface {
	Header() []byte
	Data() []byte
	RelationID() relation.ID
}

var (
	_ Page = StdPage{}
)

type TableRPage interface {
	Tuple(tupleIndex uint16) Tuple
	TupleCount() uint16
	FreeSpace() uint16
}

type TableWPage interface {
	TableRPage

	InsertTuple(data []byte) error
	UpdateTuple(tIndex TupleIndex, tuple []byte)
	DeleteTuple(tIndex TupleIndex)
	Page() []byte
}

type Id = uint32

const IdSize = 4

type Tag struct {
	PageId   Id
	Relation relation.ID
}

func NewPageTag(pageIndex Id, tableDef table.Definition) Tag {
	return Tag{
		PageId:   pageIndex,
		Relation: tableDef.RelationId(),
	}
}

const Size uint16 = dbsystem.PageSize
