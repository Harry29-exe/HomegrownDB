package data

import (
	"HomegrownDB/dbsystem/access/lob"
	"HomegrownDB/dbsystem/reldef/tabdef/column"
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
