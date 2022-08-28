package bdata

import (
	"HomegrownDB/dbsystem/access/lob"
	"HomegrownDB/dbsystem/schema/column"
)

type TupleToSave struct {
	Tuple           Tuple
	BgValuesToSave  []BgValueToSave
	LobValuesToSave []LobValueToSave
}

// BgValueToSave value to save in background table see documentation/storage_types.svg
type BgValueToSave struct {
	Value    []byte
	ColumnId column.Id
}

type LobValueToSave struct {
	Value []byte
	LobId lob.Id
}
