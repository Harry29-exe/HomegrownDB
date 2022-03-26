package schema

import "HomegrownDB/io"

type Tuple struct {
	CreatedByTx      uint32
	ModifiedByTx     uint32
	TxCommandCounter uint32
	Id               TupleId
	columns          []TupleColumn
}

type TupleId struct {
	PageId      uint32
	LinePointer uint16
}

type TupleColumn struct {
	IsPointer bool
	Data      []byte
}

func ParseTupleHeader(data []byte) *Tuple {
	deserializer := io.NewDeserializer(data)
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
