package relation

type ID = uint32

type Relation interface {
	RelationID() ID
}

func NewRelation(id ID) Relation {
	return &relation{id: id}
}

type relation struct {
	id ID
}

func (r *relation) RelationID() ID {
	return r.id
}
