package tuple

import (
	"HomegrownDB/io/bparse"
	"HomegrownDB/sql/schema/column"
	"HomegrownDB/sql/schema/table"
)

func ParseTuple(data []byte, table *table.Definition) *Tuple {
	tuple := parseTupleHeader(data)

}

type Tuple struct {
	CreatedByTx      uint32
	ModifiedByTx     uint32
	TxCommandCounter uint32
	Id               Id

	nullBitmap

	columns []Column
}

type Id struct {
	PageId      uint32
	LinePointer uint16
}

type Column struct {
	value       column.Value
	dataToParse []byte
	isParsed    bool
}

func parseTupleHeader(data []byte) *Tuple {
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
