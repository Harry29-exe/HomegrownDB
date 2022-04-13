package page

type RPage interface {
	Tuple(tupleIndex uint16) Tuple
	TupleCount() uint16
	FreeSpace() uint16
}

type WPage interface {
	InsertTuple(data []byte) error
}
