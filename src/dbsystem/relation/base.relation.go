package relation

import "HomegrownDB/common/bparse"

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
	ID      ID
	RelKind Kind
}

func (s *BaseRelation) RelationID() ID {
	return s.ID
}

func (s *BaseRelation) SetRelationID(id ID) {
	s.ID = id
}

func (s *BaseRelation) Kind() Kind {
	return s.RelKind
}
