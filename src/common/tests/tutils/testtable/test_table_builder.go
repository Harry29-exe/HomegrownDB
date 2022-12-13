package testtable

import (
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/schema/table"
)

type Builder struct {
	table table.Definition
}

func NewTestTableBuilder(name string) *Builder {
	return &Builder{table: table.NewDefinition(name)}
}

func (ttb *Builder) AddColumn(name string, nullable bool, cType ctype.Type, args ctype.Args) *Builder {
	col, err := column.NewDefinition(name, nullable, cType, args)
	if err != nil {
		panic("could not add column to table during tests")
	} else if err = ttb.table.AddColumn(col); err != nil {
		panic("could not add column to table during tests")
	}

	return ttb
}

func (ttb *Builder) SetIds(tableId table.Id, objectId relation.ID) *Builder {
	ttb.table.SetTableId(tableId)
	ttb.table.SetRelationId(objectId)

	return ttb
}

func (ttb *Builder) GetTable() table.Definition {
	table := ttb.table
	ttb.table = nil

	return table
}
