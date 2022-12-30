package pageio

import (
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/storage/dbfs"
)

func LoadStore(file dbfs.FileLike) (*Store, error) {
	//todo implement me
	panic("Not implemented")
}

func NewStore() *Store {
	return &Store{ioMap: map[relation.ID]IO{}}
}

type Store struct {
	ioMap map[relation.ID]IO
}

func (s *Store) Get(id relation.ID) IO {
	return s.ioMap[id]
}

func (s *Store) Exist(id relation.ID) bool {
	_, ok := s.ioMap[id]
	return ok
}

func (s *Store) Register(id relation.ID, io IO) {
	_, ok := s.ioMap[id]
	if ok {
		panic("Can't register io when io with same relation id is already registerd")
	}

	s.ioMap[id] = io
}
