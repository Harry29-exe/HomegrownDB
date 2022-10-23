package pageio

import (
	"HomegrownDB/dbsystem/db"
	"HomegrownDB/dbsystem/storage/dbfs"
)

var DBPageIOStore = NewStore()

func LoadStore(file dbfs.FileLike) (*Store, error) {
	//todo implement me
	panic("Not implemented")
}

func NewStore() *Store {
	return &Store{ioMap: map[db.RelationID]IO{}}
}

type Store struct {
	ioMap map[db.RelationID]IO
}

func (s *Store) Get(id db.RelationID) IO {
	return s.ioMap[id]
}

func (s *Store) Exist(id db.RelationID) bool {
	_, ok := s.ioMap[id]
	return ok
}

func (s *Store) Register(id db.RelationID, io IO) {
	_, ok := s.ioMap[id]
	if ok {
		panic("Can't register io when io with same relation id is already registerd")
	}

	s.ioMap[id] = io
}
