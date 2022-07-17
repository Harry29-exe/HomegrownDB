package tutils

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/factory"
	"HomegrownDB/dbsystem/schema/table"
)

type TestTableBuilder struct {
	table table.WDefinition
}

func NewTestTableBuilder(name string) *TestTableBuilder {
	return &TestTableBuilder{table: table.NewDefinition(name)}
}

func (ttb *TestTableBuilder) AddColumn(args column.Args) *TestTableBuilder {
	err := ttb.table.AddColumn(factory.CreateDefinition(args))
	if err != nil {
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
