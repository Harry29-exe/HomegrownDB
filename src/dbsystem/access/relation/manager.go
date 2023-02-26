package relation

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/tx"
	"log"
	"sync"
)

type Manager interface {
	Create(relation reldef.Relation, tx tx.Tx) (reldef.Relation, error)
	Delete(relation reldef.Relation) error
	AccessMngr
}

type AccessMngr interface {
	// FindByName finds relation with provided name and returns its oid,
	// if relation does not exist returns dbobj.InvalidId
	FindByName(name string) reldef.OID
	Access(oid reldef.OID, mode LockMode) reldef.Relation
	Free(relationOID reldef.OID, mode LockMode)
}

type LockMode uint8

const (
	LockNone LockMode = iota
	LockRead
	LockWrite
)

type OIDSequence interface {
	Next() reldef.OID
}

type stdManager struct {
	Buffer buffer.SharedBuffer
	FS     dbfs.FS

	mngLock *sync.RWMutex

	nameMap map[string]reldef.OID
	cache   cache
}

var _ Manager = &stdManager{}

func (s *stdManager) Delete(relation reldef.Relation) error {
	//TODO implement me
	panic("implement me")
}

func (s *stdManager) FindByName(name string) reldef.OID {
	oid, ok := s.nameMap[name]
	if !ok {
		return reldef.InvalidRelId
	}
	return oid
}

func (s *stdManager) Access(oid reldef.OID, mode LockMode) reldef.Relation {
	log.Printf("waring: locking relations is not done")
	return s.cache.relations[oid] //todo locking
}

func (s *stdManager) Free(relationOID reldef.OID, mode LockMode) {
	//todo locking
	log.Printf("waring: locking relations is not done")
}
