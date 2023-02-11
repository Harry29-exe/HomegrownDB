package pageio

import (
	"HomegrownDB/dbsystem/dbobj"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/dbfs"
	"log"
)

type Store interface {
	GetOrLoad(id dbobj.OID) IO
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

func (s *StdStore) GetOrLoad(id reldef.OID) IO {
	io, ok := s.ioMap[id]
	if !ok {
		err := s.Load(id)
		if err != nil {
			log.Panic(err.Error())
		}
		return s.ioMap[id]
	} else {
		return io
	}
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
