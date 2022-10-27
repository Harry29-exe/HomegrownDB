package page

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/schema/relation"
)

type Page interface {
	Header() []byte
	Data() []byte
	RelationID() relation.ID
}

var (
	_ Page = GenericPage{}
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

const Size uint16 = dbsystem.PageSize
