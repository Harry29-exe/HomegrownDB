package hg

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/access/relation/dbobj"
	"HomegrownDB/dbsystem/access/relation/table"
	"HomegrownDB/dbsystem/hg/di"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
)

type DBStore interface {
	TableStore() table.Store
	FS() dbfs.FS
	FsmStore() fsm.Store
	PageIOStore() pageio.Store
	SharedBuffer() buffer.SharedBuffer
	TxManager() tx.Manager

	ExecutionContainer() di.ExecutionContainer
	FrontendContainer() di.FrontendContainer

	NextRelId() relation.OID
	NextOID() dbobj.OID
}

type DB interface {
	DBStore
	RelationsOperations
	//io.Closer
	Destroy() error
}

type RelationsOperations interface {
	CreateRel(rel relation.Relation) error
	DeleteRel(rel relation.Relation) error
}
