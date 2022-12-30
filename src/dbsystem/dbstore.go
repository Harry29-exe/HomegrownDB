package dbsystem

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/dbobj"
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
)

type DBSystem interface {
	TableStore() table.Store
	FsmStore() fsm.Store
	PageIOStore() *pageio.Store
	Buffer() buffer.SharedBuffer

	NextRelId() relation.ID
	NextOID() dbobj.OID
}

var _ DBSystem = &StdSystem{}

func (s *StdSystem) Init(nextRID relation.ID, nextOID dbobj.OID) {
	s.ridCounter = appsync.NewSyncCounter(nextRID)
	s.oidCounter = appsync.NewSyncCounter(nextOID)
}

type StdSystem struct {
	Tables   table.Store
	FSMs     fsm.Store
	PageIO   *pageio.Store
	DBBuffer buffer.SharedBuffer

	ridCounter appsync.SyncCounter[relation.ID]
	oidCounter appsync.SyncCounter[dbobj.OID]
}

func (s *StdSystem) TableStore() table.Store {
	return s.Tables
}

func (s *StdSystem) FsmStore() fsm.Store {
	return s.FSMs
}

func (s *StdSystem) PageIOStore() *pageio.Store {
	return s.PageIO
}

func (s *StdSystem) Buffer() buffer.SharedBuffer {
	return s.DBBuffer
}

func (s *StdSystem) NextRelId() relation.ID {
	return s.ridCounter.GetAndIncr()
}

func (s *StdSystem) NextOID() dbobj.OID {
	return s.oidCounter.GetAndIncr()
}
