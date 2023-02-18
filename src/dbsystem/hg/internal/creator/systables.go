package creator

import (
	"HomegrownDB/dbsystem/access/systable"
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"log"
)

func createSysTables(fs dbfs.FS) error {
	creator := sysTablesCreator{FS: fs}
	relationsTable := systable.RelationsTableDef()
	columnsTable := systable.ColumnsTableDef()
	creator.cache = newPageCache(creator.FS, relationsTable, columnsTable)

	creatorTX := tx.StdTx{Id: 0}
	return creator.
		createTable(systable.HGRelationsOID, systable.HGRelationsFsmOID, systable.HGRelationsVmOID).
		createTable(systable.HGColumnsOID, systable.HGColumnsFsmOID, systable.HGColumnsVmOID).
		insertTuples(relationsTable.OID(),
			panicOnErr(systable.RelationsOps.TableAsRelationsRow(relationsTable, creatorTX)),
			panicOnErr(systable.RelationsOps.TableAsRelationsRow(columnsTable, creatorTX)),
		).
		insertTuples(systable.HGColumnsOID,
			systable.ColumnsOps.DataToRows(systable.HGRelationsOID, relationsTable.Columns(), creatorTX)...,
		).
		insertTuples(systable.HGColumnsOID,
			systable.ColumnsOps.DataToRows(systable.HGColumnsOID, columnsTable.Columns(), creatorTX)...,
		).
		flushPages().
		getError()
}

type sysTablesCreator struct {
	FS    dbfs.FS
	cache *pageCache
	err   error
}

func (c *sysTablesCreator) createTable(oid hglib.OID, fsmOID hglib.OID, vmOID hglib.OID) *sysTablesCreator {
	if c.hasErr() {
		return c
	}

	err := c.FS.InitNewPageObjectDir(oid)
	if err != nil {
		return c.error(err)
	}
	err = c.FS.InitNewPageObjectDir(fsmOID)
	if err != nil {
		return c.error(err)
	}
	err = fsm.InitFreeSpaceMapFile(fsmOID, c.FS)
	if err != nil {
		return c.error(err)
	}
	err = c.FS.InitNewPageObjectDir(vmOID)
	if err != nil {
		return c.error(err)
	}
	return c
}

func (c *sysTablesCreator) insertTuples(tableOID hglib.OID, tuples ...page.WTuple) *sysTablesCreator {
	if c.hasErr() {
		return c
	}

	for _, tuple := range tuples {
		err := c.cache.insert(tableOID, tuple)
		if err != nil {
			return c.error(err)
		}
	}
	return c
}

func (c *sysTablesCreator) flushPages() *sysTablesCreator {
	if c.hasErr() {
		return c
	}

	if err := c.cache.flush(); err != nil {
		return c.error(err)
	}
	return c
}

func (c *sysTablesCreator) error(err error) *sysTablesCreator {
	log.Println(err.Error())
	c.err = err
	return c
}

func (c *sysTablesCreator) hasErr() bool {
	return c.err != nil
}

func (c *sysTablesCreator) getError() error {
	return c.err
}

func panicOnErr[T any](val T, err error) T {
	if err != nil {
		log.Panic(err.Error())
	}
	return val
}
