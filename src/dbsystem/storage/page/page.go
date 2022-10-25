package page

import (
	"HomegrownDB/dbsystem/schema/relation"
)

func NewStdPage(bytes []byte, relID relation.ID, headerSize uint16) StdPage {
	return StdPage{
		headerSize: headerSize,
		relId:      relID,
		bytes:      bytes,
	}
}

type StdPage struct {
	headerSize uint16
	relId      relation.ID
	bytes      []byte
}

func (p StdPage) Header() []byte {
	return p.bytes[:p.headerSize]
}

func (p StdPage) Data() []byte {
	return p.bytes[p.headerSize:]
}

func (p StdPage) RelationID() relation.ID {
	return p.relId
}
