package relation

import "math"

type ID uint32

var (
	InvalidRelId ID = math.MaxUint32
	MaxRelId     ID = math.MaxUint32 - 1
)

type Relation interface {
	RelationID() ID
	Data() Data
}

type Type = uint8

const (
	TypeTable Type = iota
	TypeFsm
	TypeVm
	TypeIndex
)
