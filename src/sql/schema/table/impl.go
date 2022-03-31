package table

import (
	"HomegrownDB/io/bparse"
	"HomegrownDB/sql/schema/column"
)

type DBTable struct {
	objectId     uint64
	colNameIdMap map[string]ColumnId
	columns      map[ColumnId]column.Definition
	name         string
	byteLen      uint32
}

type tableColumn struct {
	column       column.Definition
	orderInTable ColumnId
}

func (t *DBTable) ObjectId() uint64 {
	return t.objectId
}

func (t *DBTable) Serialize() []byte {
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

func (t *DBTable) Deserialize(tableDef []byte) {
	//TODO implement me
	panic("implement me")
}

func (t *DBTable) ColumnId(name string) ColumnId {
	return t.colNameIdMap[name]
}

func (t *DBTable) ColumnsIds(names []string) []ColumnId {
	colIds := make([]ColumnId, 0, len(names))
	for i, name := range names {
		colIds[i] = t.colNameIdMap[name]
	}

	return colIds
}

func (t *DBTable) ColumnParsers(ids []ColumnId) []column.DataParser {
	parsers := make([]column.DataParser, 0, len(ids))
	for i, id := range ids {
		parsers[i] = t.columns[id].DataParser()
	}

	return parsers
}

func (t *DBTable) ColumnSerializers(ids []ColumnId) []column.DataSerializer {
	serializers := make([]column.DataSerializer, 0, len(ids))
	for i, id := range ids {
		serializers[i] = t.columns[id].DataSerializer()
	}

	return serializers
}

func (t *DBTable) AddColumn(definition column.Definition) error {
	//TODO implement me
	panic("implement me")
}

func (t *DBTable) RemoveColumn(name string) error {
	//TODO implement me
	panic("implement me")
}
