package data

import (
	"HomegrownDB/dbsystem/reldef/tabdef/column"
	"HomegrownDB/dbsystem/storage/lob"
)

type TupleToSave struct {
	Tuple           Tuple
	BgValuesToSave  []BgValueToSave
	LobValuesToSave []LobValueToSave
}

// BgValueToSave value to save in background tabdef see documentation/storage_types.svg
type BgValueToSave struct {
	Value    []byte
	ColumnId column.Id
}

type LobValueToSave struct {
	Value []byte
	LobId lob.Id
}
