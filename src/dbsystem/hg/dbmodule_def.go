package hg

import (
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/access/transaction"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/storage"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
)

type DBModule interface {
	Shutdown() error

	StorageModule() storage.Module
	ConfigModule() config.Module
	AccessModule() access.Module

	RelationManager() relation.Manager
	FS() dbfs.FS
	FsmStore() fsm.Store
	PageIOStore() pageio.Store
	SharedBuffer() buffer.SharedBuffer
	TxManager() transaction.Manager

	ExecutionContainer() ExecutionContainer
	FrontendContainer() FrontendContainer

	NextOID() hglib.OID
}

type DB interface {
	//io.Closer
	Destroy() error
	DBModule
}
