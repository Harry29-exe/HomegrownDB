package hg

import (
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

type DBModule interface {
	StorageModule() storage.Module
	ConfigModule() config.Module
	AccessModule() access.Module
	
	RelationManager() relation.Manager
	FS() dbfs.FS
	FsmStore() fsm.Store
	PageIOStore() pageio.Store
	SharedBuffer() buffer.SharedBuffer
	TxManager() tx.Manager

	ExecutionContainer() ExecutionContainer
	FrontendContainer() FrontendContainer

	NextOID() dbobj.OID
}

type DB interface {
	DBModule
	//io.Closer
	Destroy() error
}
