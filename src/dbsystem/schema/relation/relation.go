package relation

import "HomegrownDB/common/bparse"

func NewStdRelation(id ID, relKind Kind, dirFilePath string, dataSize int64) Relation {
	return &AbstractRelation{
		ID:      id,
		RelData: *NewData(dirFilePath, dataSize),
		RelKind: relKind,
	}
}

type AbstractRelation struct {
	ID      ID
	RelData data
	RelKind Kind
}

func (s *AbstractRelation) RelationID() ID {
	return s.ID
}

func (s *AbstractRelation) SetRelationID(id ID) {
	s.ID = id
}

func (s *AbstractRelation) Data() Data {
	return &s.RelData
}

func (s *AbstractRelation) Kind() Kind {
	return s.RelKind
}

func (s *AbstractRelation) Serialize(serializer *bparse.Serializer) {
	serializer.Uint32(uint32(s.ID))
	serializer.Append(s.RelData.serialize())
	serializer.Uint8(uint8(s.RelKind))
}

func (s *AbstractRelation) Deserialize(deserializer *bparse.Deserializer) {
	s.ID = ID(deserializer.Uint32())
	s.RelData = deserializeData(deserializer)
	s.RelKind = Kind(deserializer.Uint8())
}

var _ Relation = &AbstractRelation{}

func DeserializeAbstractRelation(deserializer *bparse.Deserializer) AbstractRelation {
	return AbstractRelation{
		ID:      ID(deserializer.Uint32()),
		RelData: deserializeData(deserializer),
		RelKind: Kind(deserializer.Uint8()),
	}
}
