package table

import (
	"HomegrownDB/io/bparse"
	"HomegrownDB/sql/schema/column"
	"HomegrownDB/sql/schema/column/factory"
	"errors"
)

type table struct {
	objectId     uint64
	colNameIdMap map[string]ColumnId
	columns      map[ColumnId]column.Definition
	columnsCount uint16
	name         string
}

func (t *table) ObjectId() uint64 {
	return t.objectId
}

func (t *table) Serialize() []byte {
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

func (t *table) Deserialize(tableDef []byte) {
	deserializer := bparse.NewDeserializer(tableDef)
	t.objectId = deserializer.Uint64()
	t.name = deserializer.MdString()
	t.columnsCount = deserializer.Uint16()

	data := deserializer.RemainedData()
	for i := ColumnId(0); i < t.columnsCount; i++ {
		t.columns[i], data = factory.DeserializeColumnDefinition(data)
	}
}

func (t *table) ColumnId(name string) ColumnId {
	return t.colNameIdMap[name]
}

func (t *table) ColumnsIds(names []string) []ColumnId {
	colIds := make([]ColumnId, 0, len(names))
	for i, name := range names {
		colIds[i] = t.colNameIdMap[name]
	}

	return colIds
}

func (t *table) ColumnParsers(ids []ColumnId) []column.DataParser {
	parsers := make([]column.DataParser, 0, len(ids))
	for i, id := range ids {
		parsers[i] = t.columns[id].DataParser()
	}

	return parsers
}

func (t *table) ColumnSerializers(ids []ColumnId) []column.DataSerializer {
	serializers := make([]column.DataSerializer, 0, len(ids))
	for i, id := range ids {
		serializers[i] = t.columns[id].DataSerializer()
	}

	return serializers
}

func (t *table) AddColumn(definition column.Definition) error {
	_, ok := t.colNameIdMap[definition.Name()]
	if ok {
		return errors.New("table already contains column with name:" + definition.Name())
	}
	t.columns[t.columnsCount] = definition
	t.colNameIdMap[definition.Name()] = t.columnsCount
	t.columnsCount++

	return nil
}

func (t *table) RemoveColumn(name string) error {
	colToRemoveId, ok := t.colNameIdMap[name]
	if !ok {
		return errors.New("column does not contain column with name: " + name)
	}

	newNameColIdMap := map[string]ColumnId{}
	for colName, colId := range t.colNameIdMap {
		if colId < colToRemoveId {
			newNameColIdMap[colName] = colId
		} else if colId > colToRemoveId {
			newNameColIdMap[colName] = colId - 1
		}
	}
	t.colNameIdMap = newNameColIdMap

	newColumnMap := map[ColumnId]column.Definition{}
	for colId, col := range t.columns {
		if colId < colToRemoveId {
			newColumnMap[colId] = col
		} else if colId > colToRemoveId {
			newColumnMap[colId-1] = col
		}
	}
	t.columns = newColumnMap
	t.columnsCount--
	return nil
}
