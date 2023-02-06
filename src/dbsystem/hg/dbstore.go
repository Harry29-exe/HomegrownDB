package hg

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/dbobj"
	"HomegrownDB/dbsystem/hg/di"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
)

type DBModule interface {
	RelationManager() relation.Manager
	FS() dbfs.FS
	FsmStore() fsm.Store
	PageIOStore() pageio.Store
	SharedBuffer() buffer.SharedBuffer
	TxManager() tx.Manager

	ExecutionContainer() di.ExecutionContainer
	FrontendContainer() di.FrontendContainer

	NextOID() dbobj.OID
}

type DB interface {
	DBModule
	RelationsOperations
	//io.Closer
	Destroy() error
}

type RelationsOperations interface {
	CreateRel(rel reldef.Relation) error
	DeleteRel(rel reldef.Relation) error
}
