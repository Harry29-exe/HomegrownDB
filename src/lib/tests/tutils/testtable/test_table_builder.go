package testtable

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/reldef"
)

type Builder struct {
	table        reldef.TableDefinition
	NexColtOrder reldef.Order
	NextOID      hglib.OID
}

func NewTestTableBuilder(name string) *Builder {
	return &Builder{table: reldef.NewTableDefinition(name)}
}

func (ttb *Builder) AddColumn(name string, nullable bool, typeData hgtype.ColumnType) *Builder {
	col := reldef.NewColumnDefinition(name, ttb.NextOID, ttb.NexColtOrder, typeData)
	ttb.NextOID++
	ttb.NexColtOrder++
	if err := ttb.table.AddColumn(col); err != nil {
		panic("could not add column to tabdef during tests")
	}

	return ttb
}

func (ttb *Builder) SetIds(tableId reldef.OID, objectId reldef.OID) *Builder {
	ttb.table.SetOID(tableId)
	ttb.table.SetOID(objectId)

	return ttb
}

func (ttb *Builder) GetTable() reldef.TableDefinition {
	table := ttb.table
	ttb.table = nil

	return table
}
