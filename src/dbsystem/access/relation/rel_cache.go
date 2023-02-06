package relation

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/reldef"
)

type Cache interface {
	// Update puts or replaced definition in cache
	// when rel is nil, relation with provided oid will be deleted
	Update(oid reldef.OID, rel reldef.Relation)
	// Get returns cached relations, if
	// relation does not exist it returns nil
	Get(oid reldef.OID) reldef.Relation
	// Reload purge cache and loads it from disc
	Reload(buffer buffer.SharedBuffer) error
}

var _ Cache = &cache{}

type cache struct {
	relations map[reldef.OID]reldef.Relation
}

func (c cache) Update(oid reldef.OID, rel reldef.Relation) {
	c.relations[oid] = rel
}

func (c cache) Get(oid reldef.OID) reldef.Relation {
	//TODO implement me
	panic("implement me")
}

func (c cache) Reload(buffer buffer.SharedBuffer) error {
	//TODO implement me
	panic("implement me")
}
