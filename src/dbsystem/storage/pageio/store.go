package pageio

import (
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/storage/dbfs"
)

type Store interface {
	Get(id relation.ID) IO
	Load(rel relation.Relation) error
}

func NewStore(fs dbfs.FS) *StdStore {
	return &StdStore{
		FS:    fs,
		ioMap: map[relation.ID]IO{},
	}
}

type StdStore struct {
	FS    dbfs.FS
	ioMap map[relation.ID]IO
}

func (s *StdStore) Get(id relation.ID) IO {
	return s.ioMap[id]
}

func (s *StdStore) Register(id relation.ID, io IO) {
	_, ok := s.ioMap[id]
	if ok {
		panic("Can't register io when io with same relation id is already registerd")
	}

	s.ioMap[id] = io
}

func (s *StdStore) Load(rel relation.Relation) error {
	file, err := s.FS.OpenRelationDataFile(rel)
	if err != nil {
		return err
	}
	io, err := NewPageIO(file)
	if err != nil {
		return err
	}
	s.ioMap[rel.RelationID()] = io
	return nil
}
