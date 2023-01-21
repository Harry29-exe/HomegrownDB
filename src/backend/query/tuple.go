package query

import "HomegrownDB/dbsystem/storage/tpage"

type TupleSlot struct {
	Type TupleType
}

type TupleType = uint8

const (
	TupleTypeTable TupleType = iota
	TupleTypeMem
)

type TPageTuple struct {
	TupleSlot
	Tuple tpage.Tuple
	TID   tpage.TID
}

type MemTuple struct {
	TupleSlot
}
