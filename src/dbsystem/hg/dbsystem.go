package hg

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/hg/di"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
)

var _ DBStore = &DBSystem{}

func NewDB(container *di.Container) DB {
	return &DBSystem{
		DIC:        container,
		ridCounter: appsync.NewSyncCounter(container.DBProps.NextRID),
		oidCounter: appsync.NewSyncCounter(container.DBProps.NextOID),
	}
}

type State struct {
	NextRID relation.ID
	NextOID dbobj.OID
}

type DBSystem struct {
	DIC *di.Container

	ridCounter appsync.SyncCounter[relation.ID]
	oidCounter appsync.SyncCounter[dbobj.OID]
}

func (db *DBSystem) ExecutionContainer() di.ExecutionContainer {
	return db.DIC.CreateExecutionContainer()
}

func (db *DBSystem) FrontendContainer() di.FrontendContainer {
	return db.DIC.CreateFrontendContainer()
}

func (db *DBSystem) Destroy() error {
	return db.FS().DestroyDB()
}

func (db *DBSystem) TableStore() table.Store {
	return db.DIC.TableStore
}

func (db *DBSystem) FsmStore() fsm.Store {
	return db.DIC.FsmStore
}

func (db *DBSystem) PageIOStore() pageio.Store {
	return db.DIC.PageIOStore
}

func (db *DBSystem) SharedBuffer() buffer.SharedBuffer {
	return db.DIC.SharedBuffer
}

func (db *DBSystem) FS() dbfs.FS {
	return db.DIC.FS
}

func (db *DBSystem) TxManager() tx.Manager {
	return db.DIC.TxManager
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
