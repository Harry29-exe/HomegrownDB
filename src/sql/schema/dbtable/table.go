package dbtable

import (
	. "HomegrownDB/sql/schema"
)

const TableImplName = "DbTable"

type DbTable struct {
	objectId uint64
	columns  map[string]*Column
	colList  []*Column
	name     string
	byteLen  uint32
}

func (t *DbTable) GetColumn(name string) Column {
	return *t.columns[name]
}

func (t *DbTable) ParseRow(row []byte) ParsedRow {
	for column := range t.columns {

	}
}

func parseColumn(column *Column, row []byte) []byte {
	colType := column.Type
	if colType.IsFixedSize {

	}
}

func (t *DbTable) ParseColumn(columnId uint16, row []byte) []byte {
	//TODO implement me
	panic("implement me")
}

func (t *DbTable) ParseColumns(columnsIDs []uint16, row []byte) [][]byte {
	//TODO implement me
	panic("implement me")
}
