package pageio

import (
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/storage/dbfs"
)

var DBPageIOStore = NewStore()

func LoadStore(file dbfs.FileLike) (*Store, error) {
	//todo implement me
	panic("Not implemented")
}

func NewStore() *Store {
	return &Store{ioMap: map[schema.ID]IO{}}
}

type Store struct {
	ioMap map[schema.ID]IO
}

func (s *Store) Get(id schema.ID) IO {
	return s.ioMap[id]
}

func (s *Store) Exist(id schema.ID) bool {
	_, ok := s.ioMap[id]
	return ok
}

func (s *Store) Register(id schema.ID, io IO) {
	_, ok := s.ioMap[id]
	if ok {
		panic("Can't register io when io with same relation id is already registerd")
	}

	s.ioMap[id] = io
}
