package pageio

import (
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/storage/dbfs"
)

type Store interface {
	Get(id relation.OID) IO
	Load(rel relation.Relation) error
	Register(id dbobj.OID, io IO)
}

func NewStore(fs dbfs.FS) *StdStore {
	return &StdStore{
		FS:    fs,
		ioMap: map[relation.OID]IO{},
	}
}

type StdStore struct {
	FS    dbfs.FS
	ioMap map[relation.OID]IO
}

func (s *StdStore) Get(id relation.OID) IO {
	return s.ioMap[id]
}

func (s *StdStore) Register(id dbobj.OID, io IO) {
	_, ok := s.ioMap[id]
	if ok {
		panic("Can't register io when io with same relation id is already registerd")
	}

	s.ioMap[id] = io
}

func (s *StdStore) Load(rel relation.Relation) error {
	file, err := s.FS.OpenRelationDataFile(rel.OID())
	if err != nil {
		return err
	}
	io, err := NewPageIO(file)
	if err != nil {
		return err
	}
	s.ioMap[rel.OID()] = io
	return nil
}
