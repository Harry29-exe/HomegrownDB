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

var Adapter = adapter{}

type adapter struct{}

func (a adapter) TablePage(page []byte, table table.Definition) TablePage {
	return NewPage(table, page)
}

func (a adapter) GenericPage(page []byte, rel relation.Relation) GenericPage {
	return NewGenericPage(page, rel.RelationID(), uint16(rel.PageInfo().HeaderSize))
}
