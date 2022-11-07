package relation

type ID = uint32

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
