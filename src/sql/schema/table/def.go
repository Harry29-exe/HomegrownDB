package table

import (
	"HomegrownDB/sql/schema/column"
)

type Table interface {
	ObjectId() uint64
	// Serialize table info, so it can be saved to disc and
	// later deserialize into Table object
	Serialize() []byte
	// Deserialize overrides all table info with deserialized data
	// from provided byte slice
	Deserialize(tableDef []byte)

	ColumnId(name string) ColumnId
	ColumnsIds(names []string) []ColumnId

	ColumnParsers(ids []ColumnId) []column.DataParser
	ColumnSerializers(ids []ColumnId) []column.DataSerializer

	AddColumn(definition column.Definition) error
	RemoveColumn(name string) error
}

type ColumnId = uint16
