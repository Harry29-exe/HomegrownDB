package table

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/dbobj"
	"HomegrownDB/dbsystem/schema/relation"
)

type RDefinition interface {
	relation.Relation
	bparse.Serializable

	RelationID() Id
	Name() string
	Hash() string

	BitmapLen() uint16
	ColumnCount() uint16

	CTypePattern() []hgtype.TypeData

	ColumnName(columnId column.Order) string
	ColumnOrder(name string) (order column.Order, ok bool)
	ColumnId(order column.Order) dbobj.OID

	ColumnType(id column.Order) hgtype.TypeData
	ColumnByName(name string) (col column.Def, ok bool)
	ColumnById(id dbobj.OID) column.Def
	Column(index column.Order) column.Def
	Columns() []column.Def
}

type Definition interface {
	RDefinition

	SetName(name string)

	AddColumn(definition column.WDef) error
	RemoveColumn(name string) error
}

// Id of table object, 0 if id is invalid
type Id = relation.ID

func NewDefinition(name string) Definition {
	table := &StdTable{
		AbstractRelation: relation.AbstractRelation{},
		tableId:          0,
		columns:          []column.WDef{},
		rColumns:         []column.Def{},
		name:             name,

		columnName_OrderMap: map[string]column.Order{},
		columnsNames:        nil,
		columnsCount:        0,
	}
	table.initInMemoryFields()
	return table
}

func Deserialize(data []byte) Definition {
	def := &StdTable{}
	deserializer := bparse.NewDeserializer(data)
	def.Deserialize(deserializer)

	return def
}
