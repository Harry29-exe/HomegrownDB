package reldef

import (
	"HomegrownDB/dbsystem/dbobj"
)

type OID = dbobj.OID

var (
	InvalidRelId OID = dbobj.InvalidOID
	MaxRelId     OID = dbobj.MaxOID
)

type Relation interface {
	dbobj.Obj
	SetOID(id OID)
	Kind() Kind
	Name() string

	FsmOID() dbobj.OID
	VmOID() dbobj.OID
	InitRel(id OID, fsmID OID, vmID OID)
}

type Kind uint8

const (
	TypeTable Kind = iota
	TypeIndex
)
