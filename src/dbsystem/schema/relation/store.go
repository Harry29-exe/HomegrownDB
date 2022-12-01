package relation

import (
	"HomegrownDB/common/datastructs/appsync"
	"sync"
)

// todo load store from disc
var DBStore = NewStore()

type Store interface {
	Get(id ID) Relation
	Register(relation Relation)
	Remove(id ID)

	NewRelationID() ID
}

func NewStore() Store {
	return &store{
		lock:      &sync.RWMutex{},
		relations: map[ID]Relation{},
		idCounter: appsync.NewSyncCounter[ID](0),
	}
}

type store struct {
	lock      *sync.RWMutex
	relations map[ID]Relation
	idCounter appsync.SyncCounter[ID]
}

func (s *store) Get(id ID) Relation {
	s.lock.RLock()
	defer s.lock.Unlock()
	return s.relations[id]
}

func (s *store) Register(relation Relation) {
	s.lock.Lock()
	defer s.lock.Unlock()
	id := relation.RelationID()
	_, ok := s.relations[id]
	if ok {
		panic("relation with given id already exist this is programmers mistake")
	}
	s.relations[id] = relation
}

func (s *store) Remove(id ID) {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.relations[id]
	if !ok {
		panic("can not delete not registered relation, this is programmers mistake")
	}
	delete(s.relations, id)
}

func (s *store) NewRelationID() ID {
	return s.idCounter.IncrAndGet()
}
