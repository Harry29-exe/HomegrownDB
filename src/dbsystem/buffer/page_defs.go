package buffer

type RPage interface {
	Tuple(tupleIndex uint16) Tuple
	TupleCount() uint16
	FreeSpace() uint16
}

type WPage interface {
	RPage

	InsertTuple(data []byte)
	UpdateTuple(tIndex TupleIndex, tuple []byte)
	DeleteTuple(tIndex TupleIndex)
}

type PageId = uint32

const PageIdSize = 4

const PageSize uint16 = 8192
