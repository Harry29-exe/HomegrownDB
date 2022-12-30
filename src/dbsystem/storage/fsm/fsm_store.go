package fsm

import (
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/schema/table"
)

type Store interface {
	GetFsmFor(relation relation.ID) *FreeSpaceMap
	GetFSM(fsmId relation.ID) *FreeSpaceMap
	RegisterFSM(fsm *FreeSpaceMap)
	DeleteFSM(id table.Id)
}

func NewStore() Store {
	return &StdStore{
		fsmMap:       map[relation.ID]*FreeSpaceMap{},
		parentFsmMap: map[relation.ID]*FreeSpaceMap{},
	}
}

var _ Store = &StdStore{}

type StdStore struct {
	fsmMap       map[relation.ID]*FreeSpaceMap
	parentFsmMap map[relation.ID]*FreeSpaceMap
}

func (s StdStore) GetFsmFor(relation relation.ID) *FreeSpaceMap {
	return s.parentFsmMap[relation]
}

func (s StdStore) GetFSM(fsmId table.Id) *FreeSpaceMap {
	return s.fsmMap[fsmId]
}

func (s StdStore) RegisterFSM(fsm *FreeSpaceMap) {
	s.fsmMap[fsm.RelationID()] = fsm
	s.parentFsmMap[fsm.parentRelationId] = fsm
}

func (s StdStore) DeleteFSM(id table.Id) {
	//TODO implement me
	panic("implement me")
}
