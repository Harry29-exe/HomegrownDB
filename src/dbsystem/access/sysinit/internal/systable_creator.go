package internal

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/page"
	"log"
)

type SysTablesCreator struct {
	FS    dbfs.FS
	Cache *pageCache
	err   error
}

func (c *SysTablesCreator) CreateTable(oid hglib.OID, fsmOID hglib.OID, vmOID hglib.OID) *SysTablesCreator {
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

func (c *SysTablesCreator) InsertTuples(tableOID hglib.OID, tuples ...page.WTuple) *SysTablesCreator {
	if c.hasErr() {
		return c
	}

	for _, tuple := range tuples {
		err := c.Cache.insert(tableOID, tuple)
		if err != nil {
			return c.error(err)
		}
	}
	return c
}

func (c *SysTablesCreator) InsertTuplesArr(tableOID hglib.OID, tuples ...[]page.WTuple) *SysTablesCreator {
	for _, tupleArray := range tuples {
		c.InsertTuples(tableOID, tupleArray...)
	}
	return c
}

func (c *SysTablesCreator) FlushPages() *SysTablesCreator {
	if c.hasErr() {
		return c
	}

	if err := c.Cache.flush(); err != nil {
		return c.error(err)
	}
	return c
}

func (c *SysTablesCreator) error(err error) *SysTablesCreator {
	log.Println(err.Error())
	c.err = err
	return c
}

func (c *SysTablesCreator) hasErr() bool {
	return c.err != nil
}

func (c *SysTablesCreator) GetError() error {
	return c.err
}

func PanicOnErr[T any](val T, err error) T {
	if err != nil {
		log.Panic(err.Error())
	}
	return val
}
