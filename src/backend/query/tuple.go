package query

import (
	"HomegrownDB/dbsystem/storage/page"
)

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
	Tuple page.Tuple
	TID   page.TID
}

type MemTuple struct {
	TupleSlot
}
