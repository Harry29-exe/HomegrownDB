package tabdef

import (
	"HomegrownDB/dbsystem/dbobj"
	"HomegrownDB/dbsystem/hgtype"
	reldef "HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/reldef/tabdef/column"
)

type RDefinition interface {
	reldef.Relation

	Name() string
	Hash() string

	BitmapLen() uint16
	ColumnCount() uint16

	CTypePattern() []hgtype.ColType

	ColumnName(columnId column.Order) string
	ColumnOrder(name string) (order column.Order, ok bool)
	ColumnId(order column.Order) dbobj.OID

	ColumnType(id column.Order) hgtype.ColType
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

// Id of tabdef object, 0 if id is invalid
type Id = reldef.OID

func NewDefinition(name string) Definition {
	table := &StdTable{
		BaseRelation: reldef.BaseRelation{
			RelKind: reldef.TypeTable,
		},
		columns:  []column.WDef{},
		rColumns: []column.Def{},
		name:     name,

		columnName_OrderMap: map[string]column.Order{},
		columnsNames:        nil,
		columnsCount:        0,
	}
	table.initInMemoryFields()
	return table
}
