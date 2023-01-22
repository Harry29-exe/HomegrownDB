package data

import (
	page "HomegrownDB/dbsystem/storage/page/internal"
)

type RPage interface {
	Tuple(tupleIndex uint16) Tuple
	TupleCount() uint16
	FreeSpace() uint16
	CopyBytes(dest []byte)

	PageTag() page.PageTag
}

type WPage interface {
	RPage

	InsertTuple(data []byte) error
	UpdateTuple(tIndex TupleIndex, tuple []byte)
	DeleteTuple(tIndex TupleIndex)
	Bytes() []byte
}
