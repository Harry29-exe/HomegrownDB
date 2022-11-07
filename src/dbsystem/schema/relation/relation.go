package relation

func NewStdRelation(id ID, dirFilePath string, dataSize int64) Relation {
	return &StdRelation{
		id:   id,
		data: NewData(dirFilePath, dataSize),
	}
}

type StdRelation struct {
	id   ID
	data Data
}

func (s *StdRelation) RelationID() ID {
	return s.id
}

func (s *StdRelation) Data() Data {
	return s.data
}

var _ Relation = &StdRelation{}
