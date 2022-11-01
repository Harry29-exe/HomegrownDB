package tpage

import "HomegrownDB/dbsystem/storage/page"

type TableRPage interface {
	Tuple(tupleIndex uint16) Tuple
	TupleCount() uint16
	FreeSpace() uint16

	PageTag() page.Tag
}

type TableWPage interface {
	TableRPage

	InsertTuple(data []byte) error
	UpdateTuple(tIndex TupleIndex, tuple []byte)
	DeleteTuple(tIndex TupleIndex)
	Bytes() []byte
}
