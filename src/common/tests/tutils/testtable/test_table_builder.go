package testtable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/relation/table/column"
)

type Builder struct {
	table        table.Definition
	NexColtOrder column.Order
	NextOID      dbobj.OID
}

func NewTestTableBuilder(name string) *Builder {
	return &Builder{table: table.NewDefinition(name)}
}

func (ttb *Builder) AddColumn(name string, nullable bool, typeData hgtype.ColumnType) *Builder {
	col := column.NewDefinition(name, ttb.NextOID, ttb.NexColtOrder, typeData)
	ttb.NextOID++
	ttb.NexColtOrder++
	if err := ttb.table.AddColumn(col); err != nil {
		panic("could not add column to table during tests")
	}

	return ttb
}

func (ttb *Builder) SetIds(tableId table.Id, objectId relation.OID) *Builder {
	ttb.table.SetOID(tableId)
	ttb.table.SetOID(objectId)

	return ttb
}

func (ttb *Builder) GetTable() table.Definition {
	table := ttb.table
	ttb.table = nil

	return table
}
