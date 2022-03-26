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

func (t *DbTable) GetColumnId(name string) ColumnId {
	return t.columns[name].Id
}

func (t *DbTable) GetColumnsIds(names []string) []ColumnId {
	colIds := make([]ColumnId, 0, len(names))
	for i, name := range names {
		colIds[i] = t.columns[name].Id
	}

	return colIds
}

func (t *DbTable) ParseRow(row []byte) Tuple {
	for column := range t.columns {

	}
}

func (t *DbTable) RetrieveColumn(columnId ColumnId, row []byte) TupleColumn {
	//TODO implement me
	panic("implement me")
}

func (t *DbTable) RetrieveColumns(columnsIDs []ColumnId, row []byte) []TupleColumn {
	//TODO implement me
	panic("implement me")
}

func parseColumn(column *Column, row []byte) *TupleColumn {
	colType := column.Type
	switch colType.LobStatus {
	case NEVER:

	}
}

func parseNonLobColumn(column *Column, row []byte) *TupleColumn {
	colType := column.Type
	if colType.LenPrefixSize == 0 {
		return &TupleColumn{
			IsPointer: false,
			Data:      row[:colType.ByteLen],
		}
	}

	colLen := row[:colType.LenPrefixSize]
	return &TupleColumn{
		IsPointer: false,
		Data:      row[colType.LenPrefixSize : colType.LenPrefixSize+colLen],
	}
}
