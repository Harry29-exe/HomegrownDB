package tpage

import (
	"HomegrownDB/dbsystem/storage/pageio"
)

type RPage interface {
	Tuple(tupleIndex uint16) Tuple
	TupleCount() uint16
	FreeSpace() uint16

	PageTag() pageio.PageTag
}

type WPage interface {
	RPage

	InsertTuple(data []byte) error
	UpdateTuple(tIndex TupleIndex, tuple []byte)
	DeleteTuple(tIndex TupleIndex)
	Bytes() []byte
}
