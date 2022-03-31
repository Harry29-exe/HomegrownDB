package rawdata

import (
	"HomegrownDB/io/bparse"
)

type Tuple struct {
	CreatedByTx      uint32
	ModifiedByTx     uint32
	TxCommandCounter uint32
	Id               TupleId
	columns          []ColumnValue
}

type TupleId struct {
	PageId      uint32
	LinePointer uint16
}

type ColumnValue interface {
	Value() any
	AsBytes() []byte
}

func ParseTupleHeader(data []byte) *Tuple {
	deserializer := bparse.NewDeserializer(data)
	return &Tuple{
		CreatedByTx:      deserializer.Uint32(),
		ModifiedByTx:     deserializer.Uint32(),
		TxCommandCounter: deserializer.Uint32(),
		Id: TupleId{
			PageId:      deserializer.Uint32(),
			LinePointer: deserializer.Uint16(),
		},
		columns: nil,
	}
}
