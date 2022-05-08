package page

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

type Id = uint32

const IdSize = 4

const Size uint16 = 8192
