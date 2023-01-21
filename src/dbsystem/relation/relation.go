package relation

import (
	"HomegrownDB/dbsystem/relation/dbobj"
	"math"
)

type ID = dbobj.OID

var (
	InvalidRelId ID = math.MaxUint32
	MaxRelId     ID = math.MaxUint32 - 1
)

type Relation interface {
	dbobj.Obj
	SetOID(id ID)
	Kind() Kind

	FsmOID() dbobj.OID
	VmOID() dbobj.OID
}

type Kind uint8

const (
	TypeTable Kind = iota
	TypeIndex
)
