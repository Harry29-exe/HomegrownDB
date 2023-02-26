package reldef

import (
	"HomegrownDB/dbsystem/hglib"
)

type OID = hglib.OID

var (
	InvalidRelId OID = hglib.InvalidOID
	MaxRelId     OID = hglib.MaxOID
)

type Relation interface {
	hglib.Obj
	SetOID(id OID)
	Kind() Kind
	Name() string

	FsmOID() hglib.OID
	VmOID() hglib.OID
	InitRel(id OID, fsmID OID, vmID OID)
}

type Kind uint8

const (
	TypeTable Kind = iota
	TypeIndex
	TypeSequence
)

func (k Kind) ToString() string {
	return []string{
		"TypeTable",
		"TypeIndex",
		"TypeSequence",
	}[k]
}
