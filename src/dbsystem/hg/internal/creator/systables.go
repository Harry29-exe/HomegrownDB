package creator

import (
	"HomegrownDB/dbsystem/relation/table"
	"log"
)

type sysTablesCreator struct {
	creatorCtx
	RelationsTable table.Definition
	ColumnsTable   table.Definition
}

func (c *sysTablesCreator) createSysTables() *sysTablesCreator {
	//todo implement me
	panic("Not implemented")
}

func (c *sysTablesCreator) createRelationsAndColumnsTables() *sysTablesCreator {
	if c.hasErr() {
		return c
	}

	err := c.FS.InitNewRelationDir(c.RelationsTable.OID())
	if err != nil {
		return c.error(err)
	}
	err = c.FS.InitNewRelationDir(c.ColumnsTable.OID())
	if err != nil {
		return c.error(err)
	}

	//todo implement me
	panic("Not implemented")
	//relationsTuple := page.NewTuple()
}

func (c *sysTablesCreator) error(err error) *sysTablesCreator {
	log.Println(err.Error())
	c.err = err
	return c
}

func (c *sysTablesCreator) hasErr() bool {
	return c.err != nil
}
