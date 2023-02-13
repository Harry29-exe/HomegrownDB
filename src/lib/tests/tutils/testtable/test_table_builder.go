package testtable

import (
	"HomegrownDB/dbsystem/dbobj"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/reldef/tabdef/column"
)

type Builder struct {
	table        tabdef.Definition
	NexColtOrder column.Order
	NextOID      dbobj.OID
}

func NewTestTableBuilder(name string) *Builder {
	return &Builder{table: tabdef.NewDefinition(name)}
}

func (ttb *Builder) AddColumn(name string, nullable bool, typeData hgtype.ColumnType) *Builder {
	col := column.NewDefinition(name, ttb.NextOID, ttb.NexColtOrder, typeData)
	ttb.NextOID++
	ttb.NexColtOrder++
	if err := ttb.table.AddColumn(col); err != nil {
		panic("could not add column to tabdef during tests")
	}

	return ttb
}

func (ttb *Builder) SetIds(tableId tabdef.Id, objectId reldef.OID) *Builder {
	ttb.table.SetOID(tableId)
	ttb.table.SetOID(objectId)

	return ttb
}

func (ttb *Builder) GetTable() tabdef.Definition {
	table := ttb.table
	ttb.table = nil

	return table
}
