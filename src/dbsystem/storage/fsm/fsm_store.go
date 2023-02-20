package fsm

import (
	"HomegrownDB/dbsystem/reldef"
)

// todo thought if this should be deleted as fsm generaly don't need to be locked
type Store interface {
	Register(fsm *FSM)

	GetFSM(fsmId reldef.OID) *FSM
	DeleteFSM(id reldef.OID)
}

func NewStore() Store {
	return &StdStore{
		fsmMap: map[reldef.OID]*FSM{},
	}
}

var _ Store = &StdStore{}

type StdStore struct {
	fsmMap map[reldef.OID]*FSM
}

func (s StdStore) Register(fsm *FSM) {
	s.fsmMap[fsm.fsmOID] = fsm
}

func (s StdStore) GetFSM(fsmId reldef.OID) *FSM {
	return s.fsmMap[fsmId]
}

func (s StdStore) DeleteFSM(id reldef.OID) {
	//TODO implement me
	panic("implement me")
}
