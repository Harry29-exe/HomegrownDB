package schema

type Tuple struct {
	CreatedByTx      uint32
	ModifiedByTx     uint32
	TxCommandCounter uint32
	Id               TupleId
	columns          TupleColumn
}

type TupleId struct {
	PageId      uint32
	LinePointer uint16
}

type TupleColumn struct {
	IsPointer bool
	Data      []byte
}
