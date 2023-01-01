package testtable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/relation/table/column"
)

type Builder struct {
	table table.Definition
}

func NewTestTableBuilder(name string) *Builder {
	return &Builder{table: table.NewDefinition(name)}
}

func (ttb *Builder) AddColumn(name string, nullable bool, typeData hgtype.TypeData) *Builder {
	col := column.NewDefinition(name, nullable, typeData)
	if err := ttb.table.AddColumn(col); err != nil {
		panic("could not add column to table during tests")
	}

	return ttb
}

func (ttb *Builder) SetIds(tableId table.Id, objectId relation.ID) *Builder {
	ttb.table.SetRelationID(tableId)
	ttb.table.SetRelationID(objectId)

	return ttb
}

func (ttb *Builder) GetTable() table.Definition {
	table := ttb.table
	ttb.table = nil

	return table
}
