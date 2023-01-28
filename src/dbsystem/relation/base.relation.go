package relation

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/relation/dbobj"
	"log"
)

func NewBaseRelation(relKind Kind) BaseRelation {
	return BaseRelation{
		relKind: relKind,
	}
}

func SerializeBaseRelation(rel *BaseRelation, serializer *bparse.Serializer) {
	serializer.Uint32(uint32(rel.id))
	serializer.Uint8(uint8(rel.relKind))
}

func DeserializeBaseRelation(deserializer *bparse.Deserializer) BaseRelation {
	return BaseRelation{
		id:      OID(deserializer.Uint32()),
		relKind: Kind(deserializer.Uint8()),
	}

}

var _ Relation = &BaseRelation{}

type BaseRelation struct {
	id               OID
	relKind          Kind
	freeSpaceMapOID  OID
	visibilityMapOID OID
}

func (s *BaseRelation) FsmOID() OID {
	return s.freeSpaceMapOID
}

func (s *BaseRelation) VmOID() OID {
	return s.visibilityMapOID
}

func (s *BaseRelation) OID() OID {
	return s.id
}

func (s *BaseRelation) InitRel(id OID, fsmID OID, vmID OID) {
	if s.id != dbobj.InvalidOID {
		log.Panic("relation can not be initialized multiple times")
	}
	s.id = id
	s.freeSpaceMapOID = fsmID
	s.visibilityMapOID = vmID
}

func (s *BaseRelation) SetOID(id OID) {
	s.id = id
}

func (s *BaseRelation) Kind() Kind {
	return s.relKind
}
