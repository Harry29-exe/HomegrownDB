package dbsystem

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
)

type DBSystem interface {
	TableStore() table.Store
	FsmStore() fsm.Store
	PageIOStore() pageio.Store
	Buffer() buffer.SharedBuffer

	NextRelId() relation.ID
	NextOID() dbobj.OID
}

var _ DBSystem = &StdSystem{}

type State struct {
	NextRID relation.ID
	NextOID dbobj.OID
}

type StdSystem struct {
	Tables   table.Store
	FSMs     fsm.Store
	PageIO   pageio.Store
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

func (s *StdSystem) PageIOStore() pageio.Store {
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

func (s *StdSystem) SetState(state State) error {
	s.ridCounter = appsync.NewSyncCounter(state.NextRID)
	s.oidCounter = appsync.NewSyncCounter(state.NextOID)
	return nil
}

func (s *StdSystem) GetCurrentState() State {
	return State{
		NextRID: s.ridCounter.Get(),
		NextOID: s.oidCounter.Get(),
	}
}
