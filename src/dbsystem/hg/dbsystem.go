package hg

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
)

var _ DBStore = &DBSystem{}

type State struct {
	NextRID relation.ID
	NextOID dbobj.OID
}

type DBSystem struct {
	Tables   table.Store
	FSMs     fsm.Store
	PageIO   pageio.Store
	DBBuffer buffer.SharedBuffer
	FS       dbfs.FS

	ridCounter appsync.SyncCounter[relation.ID]
	oidCounter appsync.SyncCounter[dbobj.OID]
}

func (db *DBSystem) TableStore() table.Store {
	return db.Tables
}

func (db *DBSystem) FsmStore() fsm.Store {
	return db.FSMs
}

func (db *DBSystem) PageIOStore() pageio.Store {
	return db.PageIO
}

func (db *DBSystem) Buffer() buffer.SharedBuffer {
	return db.DBBuffer
}

func (db *DBSystem) NextRelId() relation.ID {
	return db.ridCounter.GetAndIncr()
}

func (db *DBSystem) NextOID() dbobj.OID {
	return db.oidCounter.GetAndIncr()
}

func (db *DBSystem) SetState(state State) error {
	db.ridCounter = appsync.NewSyncCounter(state.NextRID)
	db.oidCounter = appsync.NewSyncCounter(state.NextOID)
	return nil
}

func (db *DBSystem) GetCurrentState() State {
	return State{
		NextRID: db.ridCounter.Get(),
		NextOID: db.oidCounter.Get(),
	}
}
