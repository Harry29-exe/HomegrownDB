package testtable

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/reldef/tabdef"
)

type Builder struct {
	table        tabdef.TableDefinition
	NexColtOrder tabdef.Order
	NextOID      hglib.OID
}

func NewTestTableBuilder(name string) *Builder {
	return &Builder{table: tabdef.NewTableDefinition(name)}
}

func (ttb *Builder) AddColumn(name string, nullable bool, typeData hgtype.ColumnType) *Builder {
	col := tabdef.NewColumnDefinition(name, ttb.NextOID, ttb.NexColtOrder, typeData)
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

func (ttb *Builder) GetTable() tabdef.TableDefinition {
	table := ttb.table
	ttb.table = nil

	return table
}
