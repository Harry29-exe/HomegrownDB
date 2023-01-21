package table

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/hgtype"
	relation "HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/relation/table/column"
)

type RDefinition interface {
	relation.Relation
	bparse.Serializable

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
	SetFsmOID(oid dbobj.OID)

	AddColumn(definition column.WDef) error
	RemoveColumn(name string) error
}

// Id of table object, 0 if id is invalid
type Id = relation.ID

func NewDefinition(name string) Definition {
	table := &StdTable{
		BaseRelation: relation.BaseRelation{},
		tableId:      0,
		columns:      []column.WDef{},
		rColumns:     []column.Def{},
		name:         name,

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
