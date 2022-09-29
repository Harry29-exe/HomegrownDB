package tutils

import (
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
)

type TestTableBuilder struct {
	table table.WDefinition
}

func NewTestTableBuilder(name string) *TestTableBuilder {
	return &TestTableBuilder{table: table.NewDefinition(name)}
}

func (ttb *TestTableBuilder) AddColumn(name string, nullable bool, cType ctype.Type, args ctype.Args) *TestTableBuilder {
	col, err := column.NewDefinition(name, nullable, cType, args)
	if err != nil {
		panic("could not add column to table during tests")
	} else if err = ttb.table.AddColumn(col); err != nil {
		panic("could not add column to table during tests")
	}

	return ttb
}

func (ttb *TestTableBuilder) SetIds(tableId table.Id, objectId uint64) *TestTableBuilder {
	ttb.table.SetTableId(tableId)
	ttb.table.SetObjectId(objectId)

	return ttb
}

func (ttb *TestTableBuilder) GetTable() table.WDefinition {
	table := ttb.table
	ttb.table = nil

	return table
}
