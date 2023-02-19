package tabdef

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/hgtype"
	reldef "HomegrownDB/dbsystem/reldef"
)

type TableRDefinition interface {
	reldef.Relation

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

// Id of tabdef object, 0 if id is invalid
type Id = reldef.OID

func NewTableDefinition(name string) TableDefinition {
	table := &Table{
		BaseRelation: reldef.BaseRelation{
			RelName: name,
			RelKind: reldef.TypeTable,
		},
		columns:  []ColumnDefinition{},
		rColumns: []ColumnRDefinition{},

		columnName_OrderMap: map[string]Order{},
		columnsCount:        0,
	}
	table.initInMemoryFields()
	return table
}
