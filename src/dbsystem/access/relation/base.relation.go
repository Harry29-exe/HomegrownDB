package relation

import (
	"HomegrownDB/dbsystem/access/relation/dbobj"
	"log"
)

var _ Relation = &BaseRelation{}

type BaseRelation struct {
	id               OID
	relKind          Kind
	freeSpaceMapOID  OID
	visibilityMapOID OID
}

func (s *BaseRelation) FsmOID() OID {
	return s.freeSpaceMapOID
}

func (s *BaseRelation) VmOID() OID {
	return s.visibilityMapOID
}

func (s *BaseRelation) OID() OID {
	return s.id
}

func (s *BaseRelation) InitRel(id OID, fsmID OID, vmID OID) {
	if s.id != dbobj.InvalidOID {
		log.Panic("relation can not be initialized multiple times")
	}
	s.id = id
	s.freeSpaceMapOID = fsmID
	s.visibilityMapOID = vmID
}

func (s *BaseRelation) SetOID(id OID) {
	s.id = id
}

func (s *BaseRelation) Kind() Kind {
	return s.relKind
}
