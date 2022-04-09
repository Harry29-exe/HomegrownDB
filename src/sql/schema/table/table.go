package table

import (
	"HomegrownDB/io/bparse"
	"HomegrownDB/sql/schema/column"
)

type Table struct {
	objectId     uint64
	colNameIdMap map[string]ColumnId
	columns      map[ColumnId]column.Definition
	columnsCount uint16
	name         string
	byteLen      uint32
}

type ColumnId = uint16

func (t *Table) ObjectId() uint64 {
	return t.objectId
}

func (t *Table) Serialize() []byte {
	serializer := bparse.NewSerializer()

	serializer.Uint64(t.objectId)
	serializer.MdString(t.name)

	columnCount := uint16(len(t.colNameIdMap))
	serializer.Uint16(columnCount)
	for _, col := range t.columns {
		serializer.Append(col.Serialize())
	}

	return serializer.GetBytes()
}

func (t *Table) Deserialize(tableDef []byte) {
	//TODO implement me
	panic("implement me")
}

func (t *Table) ColumnId(name string) ColumnId {
	return t.colNameIdMap[name]
}

func (t *Table) ColumnsIds(names []string) []ColumnId {
	colIds := make([]ColumnId, 0, len(names))
	for i, name := range names {
		colIds[i] = t.colNameIdMap[name]
	}

	return colIds
}

func (t *Table) ColumnParsers(ids []ColumnId) []column.DataParser {
	parsers := make([]column.DataParser, 0, len(ids))
	for i, id := range ids {
		parsers[i] = t.columns[id].DataParser()
	}

	return parsers
}

func (t *Table) ColumnSerializers(ids []ColumnId) []column.DataSerializer {
	serializers := make([]column.DataSerializer, 0, len(ids))
	for i, id := range ids {
		serializers[i] = t.columns[id].DataSerializer()
	}

	return serializers
}

func (t *Table) AddColumn(definition column.Definition) error {
	//TODO implement me
	panic("implement me")
}

func (t *Table) RemoveColumn(name string) error {
	//TODO implement me
	panic("implement me")
}
