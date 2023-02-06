package pageio

import (
	"HomegrownDB/dbsystem/dbobj"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/dbfs"
)

type Store interface {
	Get(id dbobj.OID) IO
	Load(rel dbobj.OID) error
	Register(id dbobj.OID, io IO)
}

func NewStore(fs dbfs.FS) Store {
	return &StdStore{
		FS:    fs,
		ioMap: map[reldef.OID]IO{},
	}
}

type StdStore struct {
	FS    dbfs.FS
	ioMap map[reldef.OID]IO
}

func (s *StdStore) Get(id reldef.OID) IO {
	return s.ioMap[id]
}

func (s *StdStore) Register(id dbobj.OID, io IO) {
	_, ok := s.ioMap[id]
	if ok {
		panic("Can't register io when io with same reldef id is already registerd")
	}

	s.ioMap[id] = io
}

func (s *StdStore) Load(oid dbobj.OID) error {
	file, err := s.FS.OpenPageObjectFile(oid)
	if err != nil {
		return err
	}
	io, err := NewPageIO(file)
	if err != nil {
		return err
	}
	s.ioMap[oid] = io
	return nil
}
