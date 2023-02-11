package relation

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/reldef"
	table "HomegrownDB/dbsystem/reldef/tabdef"
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

	if err := loader.loadRelations(); err != nil {
		return err
	} else if err = loader.loadColumns(); err != nil {
		return err
	}

	c.relations = map[reldef.OID]reldef.Relation{}
	return c.initSpecific(loader)
}

func (c *cache) initSpecific(loader *loaderCache) error {
	var err error
	for _, relation := range c.relations {
		switch relation.Kind() {
		case reldef.TypeTable:
			err = c.initTable(relation.(table.Definition), loader)
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

func (c *cache) initTable(table table.Definition, loader *loaderCache) error {
	tableCols := loader.columns[table.OID()]
	for _, col := range tableCols {
		err := table.AddColumn(col)
		if err != nil {
			return err
		}
	}

	if _, ok := c.relations[table.OID()]; ok {
		log.Panicf("loaded table with duplicated oids")
	}
	c.relations[table.OID()] = table
	return nil
}

var canNotReadPageErr = canNotReadPage{}

type canNotReadPage struct {
	Cause error
}

func (c canNotReadPage) Error() string {
	return fmt.Sprintf("can not read page becouse: %s", c.Cause.Error())
}
