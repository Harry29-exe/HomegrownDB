package creator

import (
	"HomegrownDB/dbsystem/access/systable"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
	"log"
)

type sysTablesCreator struct {
	FS    dbfs.FS
	cache *pageCache
	err   error
}

func createSysTables(fs dbfs.FS) error {
	creator := sysTablesCreator{FS: fs}
	relationsTable := systable.RelationsTableDef()
	columnsTable := systable.ColumnsTableDef()
	creator.cache = newPageCache(creator.FS, relationsTable, columnsTable)

	creatorTX := tx.StdTx{Id: 0}
	return creator.
		createTable(systable.RelationsOID, systable.RelationsFsmOID, systable.RelationsVmOID).
		createTable(systable.ColumnsOID, systable.ColumnsFsmOID, systable.ColumnsVmOID).
		insertTuples(relationsTable.OID(),
			systable.TableAsRelationsRow(relationsTable, creatorTX, 0),
			systable.TableAsRelationsRow(columnsTable, creatorTX, 0),
		).
		getError()
}

func (c *sysTablesCreator) createTable(oid dbobj.OID, fsmOID dbobj.OID, vmOID dbobj.OID) *sysTablesCreator {
	if c.hasErr() {
		return c
	}

	err := c.FS.InitNewPageObjectDir(oid)
	if err != nil {
		return c.error(err)
	}
	err = fsm.CreateFreeSpaceMap(fsmOID, c.FS)
	if err != nil {
		return c.error(err)
	}
	err = c.FS.InitNewPageObjectDir(vmOID)
	if err != nil {
		return c.error(err)
	}
	return c
}

func (c *sysTablesCreator) insertTuples(tableOID dbobj.OID, tuples ...page.Tuple) *sysTablesCreator {
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

func (c *sysTablesCreator) flushPages(oid dbobj.OID, pages []page.WPage) *sysTablesCreator {
	if c.hasErr() {
		return c
	}

	relationsFile, err := c.FS.OpenPageObjectFile(oid)
	if err != nil {
		return c.error(err)
	}
	relationsIO, err := pageio.NewPageIO(relationsFile)
	if err != nil {
		return c.error(err)
	}

	for pageID, pageToFlush := range pages {
		err = relationsIO.FlushPage(page.Id(pageID), pageToFlush.Bytes())
		if err != nil {
			return c.error(err)
		}
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
