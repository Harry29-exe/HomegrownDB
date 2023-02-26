package reldef

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/hgtype"
	"log"
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

	AddNewColumn(definition ColumnDefinition) error
	RemoveColumn(name string) error
}

// CreateTableDefinition returns new not initialized TableDefinition
func CreateTableDefinition(name string) TableDefinition {
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

// NewTableDefinition returns initialized TableDefinitions based on provided arguments
func NewTableDefinition(relation BaseRelation, columns []ColumnDefinition) (TableDefinition, error) {
	if relation.Kind() != TypeTable {
		log.Panicf("NewTableDefinition was invoked with relation.Kind() == %s",
			relation.Kind().ToString())
	}
	table := &Table{
		BaseRelation: relation,
		columns:      []ColumnDefinition{},
		rColumns:     []ColumnRDefinition{},

		columnName_OrderMap: map[string]Order{},
		columnsCount:        0,
	}
	table.initInMemoryFields()
	for _, col := range columns {
		err := table.AddColumn(col)
		if err != nil {
			return nil, err
		}
	}
	return table, nil
}
