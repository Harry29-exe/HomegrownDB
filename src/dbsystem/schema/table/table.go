package table

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/factory"
	"HomegrownDB/io/bparse"
	"errors"
)

type table struct {
	objectId     uint64
	tableId      Id
	colNameIdMap map[string]column.OrderId
	columns      map[column.OrderId]column.Definition
	columnsCount uint16
	name         string
}

func (t *table) TableId() Id {
	return t.tableId
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
	for i := column.OrderId(0); i < t.columnsCount; i++ {
		t.columns[i], data = factory.DeserializeColumnDefinition(data)
	}
}

// NullBitmapLen returns number of bytes in tuple that constitute null bitmap
func (t *table) NullBitmapLen() uint16 {
	return t.columnsCount / 8
}

func (t *table) ColumnId(name string) column.OrderId {
	return t.colNameIdMap[name]
}

func (t *table) ColumnsIds(names []string) []column.OrderId {
	colIds := make([]column.OrderId, 0, len(names))
	for i, name := range names {
		colIds[i] = t.colNameIdMap[name]
	}

	return colIds
}

func (t *table) ColumnParser(id column.OrderId) column.DataParser {
	return t.columns[id].DataParser()
}

func (t *table) ColumnParsers(ids []column.OrderId) []column.DataParser {
	parsers := make([]column.DataParser, 0, len(ids))
	for i, id := range ids {
		parsers[i] = t.columns[id].DataParser()
	}

	return parsers
}

func (t *table) ColumnSerializer(id column.OrderId) column.DataSerializer {
	return t.columns[id].DataSerializer()
}

func (t *table) ColumnSerializers(ids []column.OrderId) []column.DataSerializer {
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

	newNameColIdMap := map[string]column.OrderId{}
	for colName, colId := range t.colNameIdMap {
		if colId < colToRemoveId {
			newNameColIdMap[colName] = colId
		} else if colId > colToRemoveId {
			newNameColIdMap[colName] = colId - 1
		}
	}
	t.colNameIdMap = newNameColIdMap

	newColumnMap := map[column.OrderId]column.Definition{}
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
