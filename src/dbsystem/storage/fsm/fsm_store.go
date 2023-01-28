package fsm

import (
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/table"
)

// todo thought if this should be deleted as fsm generaly don't need to be locked
type Store interface {
	Register(fsm *FSM)

	GetFSM(fsmId relation.OID) *FSM
	DeleteFSM(id table.Id)
}

func NewStore() Store {
	return &StdStore{
		fsmMap: map[relation.OID]*FSM{},
	}
}

var _ Store = &StdStore{}

type StdStore struct {
	fsmMap map[relation.OID]*FSM
}

func (s StdStore) Register(fsm *FSM) {
	s.fsmMap[fsm.fsmOID] = fsm
}

func (s StdStore) GetFSM(fsmId table.Id) *FSM {
	return s.fsmMap[fsmId]
}

func (s StdStore) DeleteFSM(id table.Id) {
	//TODO implement me
	panic("implement me")
}
