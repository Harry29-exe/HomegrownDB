package dbsystem

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/dbobj"
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/fsm"
)

type DBSystem interface {
	TableStore() table.Store
	FsmStore() fsm.Store
	Buffer() buffer.SharedBuffer

	NextRelId() relation.ID
	NextOID() dbobj.OID
}

var _ DBSystem = &StdSystem{}

type StdSystem struct {
	tableStore table.Store
	fsmStore   fsm.Store
	dbBuffer   buffer.SharedBuffer

	ridCounter appsync.SyncCounter[relation.ID]
	oidCounter appsync.SyncCounter[dbobj.OID]
}

func (s *StdSystem) NextRelId() relation.ID {
	return s.ridCounter.GetAndIncr()
}

func (s *StdSystem) NextOID() dbobj.OID {
	return s.oidCounter.GetAndIncr()
}

func (s *StdSystem) TableStore() table.Store {
	return s.tableStore
}

func (s *StdSystem) FsmStore() fsm.Store {
	return s.fsmStore
}

func (s *StdSystem) Buffer() buffer.SharedBuffer {
	return s.dbBuffer
}
