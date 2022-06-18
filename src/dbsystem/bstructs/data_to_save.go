package bstructs

import (
	"HomegrownDB/dbsystem/io/bgtable"
	"HomegrownDB/dbsystem/schema/table"
	"bytes"
)

type TupleToSave struct {
	Tuple           Tuple
	BgValuesToSave  []BgValueToSave
	LobValuesToSave []LobValueToSave
}

// BgValueToSave value to save in background table see documentation/storage_types.svg
type BgValueToSave struct {
	Value []byte
	BgId  bgtable.BgId
}

type LobValueToSave struct {
	Value []byte
}

type tupleBuilder struct {
	table        table.Definition
	sortedValues []any

	buffer    bytes.Buffer
	bufferPtr InTuplePtr

	lobs     [][]byte
	bgValues [][]byte
}
