package hg

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
)

var _ DBStore = &DBSystem{}

type State struct {
	NextRID relation.ID
	NextOID dbobj.OID
}

type DBSystem struct {
	Tables   table.Store
	FSMs     fsm.Store
	PageIO   pageio.Store
	DBBuffer buffer.SharedBuffer
	FS       dbfs.FS

	ridCounter appsync.SyncCounter[relation.ID]
	oidCounter appsync.SyncCounter[dbobj.OID]
}

func (s *DBSystem) TableStore() table.Store {
	return s.Tables
}

func (s *DBSystem) FsmStore() fsm.Store {
	return s.FSMs
}

func (s *DBSystem) PageIOStore() pageio.Store {
	return s.PageIO
}

func (s *DBSystem) Buffer() buffer.SharedBuffer {
	return s.DBBuffer
}

func (s *DBSystem) NextRelId() relation.ID {
	return s.ridCounter.GetAndIncr()
}

func (s *DBSystem) NextOID() dbobj.OID {
	return s.oidCounter.GetAndIncr()
}

func (s *DBSystem) SetState(state State) error {
	s.ridCounter = appsync.NewSyncCounter(state.NextRID)
	s.oidCounter = appsync.NewSyncCounter(state.NextOID)
	return nil
}

func (s *DBSystem) GetCurrentState() State {
	return State{
		NextRID: s.ridCounter.Get(),
		NextOID: s.oidCounter.Get(),
	}
}
