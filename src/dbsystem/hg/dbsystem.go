package hg

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/dbobj"
	"HomegrownDB/dbsystem/storage"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
)

var _ DBModule = &DBSystem{}

type State struct {
	NextOID dbobj.OID
}

type DBSystem struct {
	storageModule storage.Module
	configModule  config.Module
	accessModule  access.Module

	oidCounter appsync.SyncCounter[dbobj.OID]
}

func (db *DBSystem) StorageModule() storage.Module {
	return db.storageModule
}

func (db *DBSystem) ConfigModule() config.Module {
	return db.configModule
}

func (db *DBSystem) AccessModule() access.Module {
	return db.accessModule
}

func (db *DBSystem) ExecutionContainer() ExecutionContainer {
	return ExecutionContainer{
		SharedBuffer:    db.accessModule.SharedBuffer(),
		FsmStore:        nil,
		RelationManager: db.accessModule.RelationManager(),
	}
}

func (db *DBSystem) FrontendContainer() FrontendContainer {
	return FrontendContainer{
		AuthManger:         nil,
		ExecutionContainer: db.ExecutionContainer(),
		TxManager:          nil,
	}
}

func (db *DBSystem) Destroy() error {
	return db.FS().DestroyDB()
}

func (db *DBSystem) RelationManager() relation.Manager {
	return db.accessModule.RelationManager()
}

func (db *DBSystem) FsmStore() fsm.Store {
	return nil
}

func (db *DBSystem) PageIOStore() pageio.Store {
	return db.storageModule.PageIOStore()
}

func (db *DBSystem) SharedBuffer() buffer.SharedBuffer {
	return db.accessModule.SharedBuffer()
}

func (db *DBSystem) FS() dbfs.FS {
	return db.storageModule.FS()
}

func (db *DBSystem) TxManager() tx.Manager {
	return nil
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
