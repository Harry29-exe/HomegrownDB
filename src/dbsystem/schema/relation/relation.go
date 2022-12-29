package relation

func NewStdRelation(id ID, relKind Kind, dirFilePath string, dataSize int64) Relation {
	return &StdRelation{
		id:   id,
		data: NewData(dirFilePath, dataSize),
	}
}

type StdRelation struct {
	id   ID
	data Data
	kind Kind
}

func (s *StdRelation) RelationID() ID {
	return s.id
}

func (s *StdRelation) Data() Data {
	return s.data
}

func (s *StdRelation) Kind() Kind {
	return s.kind
}

var _ Relation = &StdRelation{}
