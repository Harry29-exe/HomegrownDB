package dbtable

import (
	"HomegrownDB/sql/schema"
)

const TableImplName = "DbTable"

type DbTable struct {
	objectId uint64
	columns  map[string]*schema.Column
	colList  []*schema.Column
	name     string
	byteLen  uint32
}

func (t *DbTable) GetColumn(name string) schema.Column {
	return *t.columns[name]
}
