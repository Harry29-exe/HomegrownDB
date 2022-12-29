package dbsystem

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/fsm"
)

type DBSystem interface {
	TableStore() table.Store
	FsmStore() fsm.Store
	Buffer() buffer.SharedBuffer
}

var _ DBSystem = &StdSystem{}

type StdSystem struct {
	tableStore table.Store
	fsmStore   fsm.Store
	dbBuffer   buffer.SharedBuffer
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
