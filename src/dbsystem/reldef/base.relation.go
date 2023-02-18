package reldef

import (
	"HomegrownDB/dbsystem/hglib"
	"log"
)

var _ Relation = &BaseRelation{}

type BaseRelation struct {
	RelName          string
	ID               OID
	RelKind          Kind
	FreeSpaceMapOID  OID
	VisibilityMapOID OID
}

func (s *BaseRelation) Name() string {
	return s.RelName
}

func (s *BaseRelation) FsmOID() OID {
	return s.FreeSpaceMapOID
}

func (s *BaseRelation) VmOID() OID {
	return s.VisibilityMapOID
}

func (s *BaseRelation) OID() OID {
	return s.ID
}

func (s *BaseRelation) InitRel(id OID, fsmID OID, vmID OID) {
	if s.ID != hglib.InvalidOID {
		log.Panic("reldef can not be initialized multiple times")
	}
	s.ID = id
	s.FreeSpaceMapOID = fsmID
	s.VisibilityMapOID = vmID
}

func (s *BaseRelation) SetOID(id OID) {
	s.ID = id
}

func (s *BaseRelation) Kind() Kind {
	return s.RelKind
}
