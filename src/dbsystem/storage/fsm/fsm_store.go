package fsm

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/schema/table"
)

type Store interface {
	GetFsmFor(relation relation.ID) *FreeSpaceMap
	GetFSM(fsmId relation.ID) *FreeSpaceMap
	CreateNewFSM(fsmRelation, parentRelation relation.Relation) error
	LoadFSM(fsmRelation, parentRelation relation.Relation) error
	DeleteFSM(id table.Id)
}

func NewStore(sharedBuffer buffer.SharedBuffer) Store {
	return &StdStore{
		fsmMap:       map[relation.ID]*FreeSpaceMap{},
		parentFsmMap: map[relation.ID]*FreeSpaceMap{},
		buffer:       sharedBuffer,
	}
}

var _ Store = &StdStore{}

type StdStore struct {
	fsmMap       map[relation.ID]*FreeSpaceMap
	parentFsmMap map[relation.ID]*FreeSpaceMap
	buffer       buffer.SharedBuffer
}

func (s StdStore) GetFsmFor(relation relation.ID) *FreeSpaceMap {
	return s.parentFsmMap[relation]
}

func (s StdStore) GetFSM(fsmId table.Id) *FreeSpaceMap {
	return s.fsmMap[fsmId]
}

func (s StdStore) CreateNewFSM(fsmRelation, parentRelation relation.Relation) error {
	fsm, err := CreateFreeSpaceMap(fsmRelation, parentRelation, s.buffer)
	if err != nil {
		return err
	}
	s.fsmMap[fsmRelation.RelationID()] = fsm
	s.parentFsmMap[parentRelation.RelationID()] = fsm
	return err
}

func (s StdStore) LoadFSM(fsmRelation, parentRelation relation.Relation) error {
	fsm, err := LoadFreeSpaceMap(fsmRelation, parentRelation, s.buffer)
	if err != nil {
		return err
	}
	s.fsmMap[fsmRelation.RelationID()] = fsm
	s.parentFsmMap[parentRelation.RelationID()] = fsm
	return nil
}

func (s StdStore) DeleteFSM(id table.Id) {
	//TODO implement me
	panic("implement me")
}
