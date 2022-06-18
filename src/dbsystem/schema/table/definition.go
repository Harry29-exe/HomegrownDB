package table

import (
	"HomegrownDB/dbsystem/schema/column"
)

type Definition interface {
	TableId() Id
	ObjectId() uint64
	// Serialize table info, so it can be saved to disc and
	// later deserialize into table object
	Serialize() []byte
	// Deserialize overrides all table info with deserialized data
	// from provided byte slice
	Deserialize(tableDef []byte)
	BitmapLen() uint16
	ColumnCount() uint16

	ColumnName(columnId column.OrderId) string
	ColumnId(name string) column.OrderId
	ColumnsIds(names []string) []column.OrderId

	ColumnParser(id column.OrderId) column.DataParser
	ColumnParsers(ids []column.OrderId) []column.DataParser
	ColumnSerializer(id column.OrderId) column.DataSerializer
	ColumnSerializers(ids []column.OrderId) []column.DataSerializer
	AllColumnSerializer() []column.DataSerializer

	AddColumn(definition column.Definition) error
	RemoveColumn(name string) error
}

type Id = uint32

func NewDefinition(name string, tableId Id, objId uint64) Definition {
	return &table{
		tableId:      tableId,
		objectId:     objId,
		colNameIdMap: map[string]column.OrderId{},
		columnsNames: []string{},
		columns:      []column.Definition{},
		columnsCount: 0,
		name:         name,
	}
}
