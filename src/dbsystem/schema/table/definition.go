package table

import (
	"HomegrownDB/dbsystem/schema/column"
)

type Definition interface {
	ObjectId() uint64
	// Serialize table info, so it can be saved to disc and
	// later deserialize into table object
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

func NewDefinition(name string, objId uint64) Definition {
	return &table{
		objectId:     objId,
		colNameIdMap: map[string]ColumnId{},
		columns:      map[ColumnId]column.Definition{},
		columnsCount: 0,
		name:         name,
	}
}
