package fsm

import (
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/reldef/tabdef"
)

// todo thought if this should be deleted as fsm generaly don't need to be locked
type Store interface {
	Register(fsm *FSM)

	GetFSM(fsmId reldef.OID) *FSM
	DeleteFSM(id tabdef.Id)
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

func (s StdStore) GetFSM(fsmId tabdef.Id) *FSM {
	return s.fsmMap[fsmId]
}

func (s StdStore) DeleteFSM(id tabdef.Id) {
	//TODO implement me
	panic("implement me")
}
