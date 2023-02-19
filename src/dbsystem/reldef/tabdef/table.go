package tabdef

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/hgtype"
	reldef "HomegrownDB/dbsystem/reldef"
)

type RDefinition interface {
	reldef.Relation

	Name() string
	Hash() string

	BitmapLen() uint16
	ColumnCount() uint16

	CTypePattern() []hgtype.ColType

	ColumnName(columnId Order) string
	ColumnOrder(name string) (order Order, ok bool)
	ColumnId(order Order) hglib.OID

	ColumnType(id Order) hgtype.ColType
	ColumnByName(name string) (col ColumnRDefinition, ok bool)
	ColumnById(id hglib.OID) ColumnRDefinition
	Column(index Order) ColumnRDefinition
	Columns() []ColumnRDefinition
}

type Definition interface {
	RDefinition

	SetName(name string)

	AddColumn(definition ColumnDefinition) error
	RemoveColumn(name string) error
}

// Id of tabdef object, 0 if id is invalid
type Id = reldef.OID

func NewDefinition(name string) Definition {
	table := &StdTable{
		BaseRelation: reldef.BaseRelation{
			RelName: name,
			RelKind: reldef.TypeTable,
		},
		columns:  []ColumnDefinition{},
		rColumns: []ColumnRDefinition{},

		columnName_OrderMap: map[string]Order{},
		columnsNames:        nil,
		columnsCount:        0,
	}
	table.initInMemoryFields()
	return table
}
