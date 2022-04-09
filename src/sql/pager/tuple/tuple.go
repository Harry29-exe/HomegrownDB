package tuple

import (
	"HomegrownDB/io/bparse"
	"HomegrownDB/sql/schema/column"
)

type Tuple struct {
	CreatedByTx      uint32
	ModifiedByTx     uint32
	TxCommandCounter uint32
	Id               Id
	columns          []column.Value
}

type Id struct {
	PageId      uint32
	LinePointer uint16
}

func ParseTupleHeader(data []byte) *Tuple {
	deserializer := bparse.NewDeserializer(data)
	return &Tuple{
		CreatedByTx:      deserializer.Uint32(),
		ModifiedByTx:     deserializer.Uint32(),
		TxCommandCounter: deserializer.Uint32(),
		Id: Id{
			PageId:      deserializer.Uint32(),
			LinePointer: deserializer.Uint16(),
		},
		columns: nil,
	}
}
