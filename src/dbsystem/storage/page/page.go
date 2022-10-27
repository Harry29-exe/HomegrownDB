package page

import (
	"HomegrownDB/dbsystem/schema/relation"
)

func NewGenericPage(bytes []byte, rel relation.Relation) GenericPage {
	return GenericPage{
		headerSize: uint16(rel.PageInfo().HeaderSize),
		relId:      rel.RelationID(),
		bytes:      bytes,
	}
}

type GenericPage struct {
	headerSize uint16
	relId      relation.ID
	bytes      []byte
}

func (p GenericPage) Header() []byte {
	return p.bytes[:p.headerSize]
}

func (p GenericPage) Data() []byte {
	return p.bytes[p.headerSize:]
}

func (p GenericPage) RelationID() relation.ID {
	return p.relId
}
