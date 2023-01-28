package relation

import (
	"HomegrownDB/common/datastructs/appsync"
	"sync"
)

type Store interface {
	Get(id OID) Relation
	Register(relation Relation)
	Remove(id OID)

	NewRelationID() OID
}

func NewStore() Store {
	return &store{
		lock:      &sync.RWMutex{},
		relations: map[OID]Relation{},
		idCounter: appsync.NewSyncCounter[OID](0),
	}
}

type store struct {
	lock      *sync.RWMutex
	relations map[OID]Relation
	idCounter appsync.SyncCounter[OID]
}

func (s *store) Get(id OID) Relation {
	s.lock.RLock()
	defer s.lock.Unlock()
	return s.relations[id]
}

func (s *store) Register(relation Relation) {
	s.lock.Lock()
	defer s.lock.Unlock()
	id := relation.OID()
	_, ok := s.relations[id]
	if ok {
		panic("relation with given id already exist this is programmers mistake")
	}
	s.relations[id] = relation
}

func (s *store) Remove(id OID) {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.relations[id]
	if !ok {
		panic("can not delete not registered relation, this is programmers mistake")
	}
	delete(s.relations, id)
}

func (s *store) NewRelationID() OID {
	return s.idCounter.IncrAndGet()
}
