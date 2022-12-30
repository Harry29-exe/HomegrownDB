package relation

import "HomegrownDB/common/bparse"

func NewBaseRelation(id ID, relKind Kind, dirFilePath string, dataSize int64) BaseRelation {
	return BaseRelation{
		ID:      id,
		RelData: *NewData(dirFilePath, dataSize),
		RelKind: relKind,
	}
}

func SerializeBaseRelation(rel *BaseRelation, serializer *bparse.Serializer) {
	serializer.Uint32(uint32(rel.ID))
	serializer.Append(rel.RelData.serialize())
	serializer.Uint8(uint8(rel.RelKind))
}

func DeserializeBaseRelation(deserializer *bparse.Deserializer) BaseRelation {
	return BaseRelation{
		ID:      ID(deserializer.Uint32()),
		RelData: deserializeData(deserializer),
		RelKind: Kind(deserializer.Uint8()),
	}

}

var _ Relation = &BaseRelation{}

type BaseRelation struct {
	ID      ID
	RelData data
	RelKind Kind
}

func (s *BaseRelation) RelationID() ID {
	return s.ID
}

func (s *BaseRelation) SetRelationID(id ID) {
	s.ID = id
}

func (s *BaseRelation) Data() Data {
	return &s.RelData
}

func (s *BaseRelation) Kind() Kind {
	return s.RelKind
}
