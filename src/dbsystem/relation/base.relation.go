package relation

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/relation/dbobj"
)

func NewBaseRelation(id ID, relKind Kind, dirFilePath string, dataSize int64) BaseRelation {
	return BaseRelation{
		ID:      id,
		RelKind: relKind,
	}
}

func SerializeBaseRelation(rel *BaseRelation, serializer *bparse.Serializer) {
	serializer.Uint32(uint32(rel.ID))
	serializer.Uint8(uint8(rel.RelKind))
}

func DeserializeBaseRelation(deserializer *bparse.Deserializer) BaseRelation {
	return BaseRelation{
		ID:      ID(deserializer.Uint32()),
		RelKind: Kind(deserializer.Uint8()),
	}

}

var _ Relation = &BaseRelation{}

type BaseRelation struct {
	ID               ID
	RelKind          Kind
	FreeSpaceMapOID  dbobj.OID
	VisibilityMapOID dbobj.OID
}

func (s *BaseRelation) FsmOID() dbobj.OID {
	return s.FreeSpaceMapOID
}

func (s *BaseRelation) VmOID() dbobj.OID {
	return s.VisibilityMapOID
}

func (s *BaseRelation) OID() ID {
	return s.ID
}

func (s *BaseRelation) SetOID(id ID) {
	s.ID = id
}

func (s *BaseRelation) Kind() Kind {
	return s.RelKind
}
