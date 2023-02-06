package hg

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/dbobj"
	"HomegrownDB/dbsystem/hg/di"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
)

var _ DBModule = &DBSystem{}

func NewDB(container *di.Container) DB {
	return &DBSystem{
		DIC:        container,
		oidCounter: appsync.NewSyncCounter(container.DBProps.NextOID),
	}
}

type State struct {
	NextOID dbobj.OID
}

type DBSystem struct {
	DIC *di.Container

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

func (db *DBSystem) RelationManager() relation.Manager {
	return db.DIC.RelationManager
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

func (db *DBSystem) NextOID() dbobj.OID {
	return db.oidCounter.GetAndIncr()
}

func (db *DBSystem) SetState(state State) error {
	db.oidCounter = appsync.NewSyncCounter(state.NextOID)
	return nil
}

func (db *DBSystem) GetCurrentState() State {
	return State{
		NextOID: db.oidCounter.Get(),
	}
}
