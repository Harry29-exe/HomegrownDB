package relation

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/reldef"
	"fmt"
	"log"
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

func (c *cache) Update(oid reldef.OID, rel reldef.Relation) {
	c.relations[oid] = rel
}

func (c *cache) Get(oid reldef.OID) reldef.Relation {
	return c.relations[oid]
}

func (c *cache) Reload(buffer buffer.SharedBuffer) error {
	loader := newLoaderCache(buffer)

	if err := loader.load(); err != nil {
		return err
	}

	c.relations = map[reldef.OID]reldef.Relation{}
	return c.load(loader)
}

func (c *cache) load(loader *loaderCache) error {
	var err error
	for _, relation := range loader.relations {
		switch relation.Kind() {
		case reldef.TypeTable:
			err = c.loadTable(relation, loader)
		case reldef.TypeSequence:
			err = c.loadSequence(relation, loader)
		default:
			//todo implement me
			panic("Not implemented")
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func (c *cache) loadTable(tableRel reldef.BaseRelation, loader *loaderCache) error {
	tableCols := loader.columns[tableRel.OID()]
	table, err := reldef.NewTableDefinition(tableRel, tableCols)
	if err != nil {
		return err
	}

	if _, ok := c.relations[table.OID()]; ok {
		log.Panicf("loaded table with duplicated oids")
	}
	c.relations[table.OID()] = table
	return nil
}

func (c *cache) loadSequence(seqRel reldef.BaseRelation, loader *loaderCache) error {
	futureSeq := loader.sequences[seqRel.ID]
	sequence := futureSeq.Init(seqRel)

	if _, ok := c.relations[sequence.OID()]; ok {
		log.Panicf("loaded sequence with duplicated oids")
	}
	c.relations[sequence.OID()] = sequence
	return nil
}

// -------------------------
//      errors
// -------------------------

var canNotReadPageErr = canNotReadPage{}

type canNotReadPage struct {
	Cause error
}

func (c canNotReadPage) Error() string {
	return fmt.Sprintf("can not read page becouse: %s", c.Cause.Error())
}

func (c canNotReadPage) Is(err error) bool {
	_, ok := err.(canNotReadPage)
	return ok
}
