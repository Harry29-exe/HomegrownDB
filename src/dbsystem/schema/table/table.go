package table

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/relation"
)

type RDefinition interface {
	relation.Relation

	TableId() Id
	Name() string
	Hash() string

	// Serialize table info, so it can be saved to disc and
	// later deserialize into table object
	Serialize() []byte
	// Deserialize overrides all table info with deserialized data
	// from provided byte slice
	Deserialize(tableDef []byte)
	BitmapLen() uint16
	ColumnCount() uint16

	CTypePattern() []hgtype.Wrapper

	ColumnName(columnId column.Order) string
	ColumnOrder(name string) (order column.Order, ok bool)
	ColumnId(order column.Order) column.Id

	ColumnType(id column.Order) hgtype.Wrapper
	ColumnByName(name string) (col column.Def, ok bool)
	ColumnById(id column.Id) column.Def
	Column(index column.Order) column.Def
	Columns() []column.Def
}

type Definition interface {
	RDefinition

	SetTableId(id Id)
	SetRelationId(id relation.ID)
	SetName(name string)

	AddColumn(definition column.WDef) error
	RemoveColumn(name string) error
}

// Id of table object, 0 if id is invalid
type Id = relation.ID

func NewDefinition(name string) Definition {
	table := &StdTable{
		tableId:  0,
		objectId: 0,
		columns:  []column.WDef{},
		rColumns: []column.Def{},
		name:     name,

		nextColumnId:        appsync.NewSyncCounter[column.Id](0),
		columnName_OrderMap: map[string]column.Order{},
		columnsNames:        nil,
		columnsCount:        0,
	}
	table.initInMemoryFields()
	return table
}

func Deserialize(data []byte) Definition {
	def := &StdTable{}
	def.Deserialize(data)

	return def
}
