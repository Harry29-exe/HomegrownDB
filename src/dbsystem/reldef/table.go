package reldef

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/hgtype"
)

type TableRDefinition interface {
	Relation

	Name() string
	Hash() string

	BitmapLen() uint16
	ColumnCount() uint16

	ColumnType(id Order) hgtype.ColType
	ColumnByName(name string) (col ColumnRDefinition, ok bool)
	ColumnById(id hglib.OID) ColumnRDefinition
	Column(index Order) ColumnRDefinition
	Columns() []ColumnRDefinition
}

type TableDefinition interface {
	TableRDefinition

	SetName(name string)

	AddColumn(definition ColumnDefinition) error
	RemoveColumn(name string) error
}

func NewTableDefinition(name string) TableDefinition {
	table := &Table{
		BaseRelation: BaseRelation{
			RelName: name,
			RelKind: TypeTable,
		},
		columns:  []ColumnDefinition{},
		rColumns: []ColumnRDefinition{},

		columnName_OrderMap: map[string]Order{},
		columnsCount:        0,
	}
	table.initInMemoryFields()
	return table
}
