package hg

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
)

type DBStore interface {
	TableStore() table.Store
	FsmStore() fsm.Store
	PageIOStore() pageio.Store
	Buffer() buffer.SharedBuffer

	NextRelId() relation.ID
	NextOID() dbobj.OID
}

type DB interface {
	CreateRel(rel relation.Relation) error
}
