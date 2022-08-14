package table

import (
	"HomegrownDB/dbsystem/schema/column"
)

type Definition interface {
	TableId() Id
	ObjectId() uint64
	Name() string

	// Serialize table info, so it can be saved to disc and
	// later deserialize into table object
	Serialize() []byte
	// Deserialize overrides all table info with deserialized data
	// from provided byte slice
	Deserialize(tableDef []byte)
	BitmapLen() uint16
	ColumnCount() uint16

	ColumnName(columnId column.OrderId) string
	ColumnId(name string) (id column.OrderId, ok bool)
	ColumnsIds(names []string) []column.OrderId

	ColumnParser(id column.OrderId) column.DataParser
	ColumnParsers(ids []column.OrderId) []column.DataParser
	AllColumnParsers() []column.DataParser
	ColumnSerializer(id column.OrderId) column.DataSerializer
	ColumnSerializers(ids []column.OrderId) []column.DataSerializer
	AllColumnSerializer() []column.DataSerializer

	GetColumn(index column.OrderId) column.Definition
}

type WDefinition interface {
	Definition

	SetTableId(id Id)
	SetObjectId(id uint64)
	SetName(name string)

	AddColumn(definition column.WDefinition) error
	RemoveColumn(name string) error
}

// Id of table object, 0 means id is invalid
type Id = uint16

func NewDefinition(name string) WDefinition {
	table := &StandardTable{
		tableId:  0,
		objectId: 0,
		columns:  []column.WDefinition{},
		name:     name,

		colNameIdMap:  map[string]column.OrderId{},
		columnsNames:  nil,
		columnsCount:  0,
		columnParsers: nil,
	}
	table.initInMemoryFields()
	return table
}

func Deserialize(data []byte) WDefinition {
	def := &StandardTable{}
	def.Deserialize(data)

	return def
}
