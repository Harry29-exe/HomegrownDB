package table

import (
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
)

type Definition interface {
	TableId() Id
	// OID Object id
	OID() uint64
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

	ColumnType(id column.OrderId) ctype.CType
	ColumnByName(name string) (col column.Def, ok bool)
	Column(index column.OrderId) column.Def
	Columns() []column.Def
}

type WDefinition interface {
	Definition

	SetTableId(id Id)
	SetObjectId(id uint64)
	SetName(name string)

	AddColumn(definition column.WDef) error
	RemoveColumn(name string) error
}

// Id of table object, 0 if id is invalid
type Id = uint16

func NewDefinition(name string) WDefinition {
	table := &StandardTable{
		tableId:  0,
		objectId: 0,
		columns:  []column.WDef{},
		rColumns: []column.Def{},
		name:     name,

		colNameIdMap: map[string]column.OrderId{},
		columnsNames: nil,
		columnsCount: 0,
	}
	table.initInMemoryFields()
	return table
}

func Deserialize(data []byte) WDefinition {
	def := &StandardTable{}
	def.Deserialize(data)

	return def
}
