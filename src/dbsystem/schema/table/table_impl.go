package table

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/factory"
	"errors"
	"math"
)

type StandardTable struct {
	objectId     uint64
	tableId      Id
	colNameIdMap map[string]column.OrderId
	columnsNames []string
	columns      []column.Definition
	columnsCount uint16
	name         string
}

func (t *StandardTable) SetTableId(id Id) {
	t.tableId = id
}

func (t *StandardTable) TableId() Id {
	return t.tableId
}

func (t *StandardTable) SetObjectId(id uint64) {
	t.objectId = id
}

func (t *StandardTable) ObjectId() uint64 {
	return t.objectId
}

func (t *StandardTable) SetName(name string) {
	t.name = name
}

func (t *StandardTable) Name() string {
	return t.name
}

func (t *StandardTable) Serialize() []byte {
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

func (t *StandardTable) Deserialize(tableDef []byte) {
	deserializer := bparse.NewDeserializer(tableDef)
	t.objectId = deserializer.Uint64()
	t.name = deserializer.MdString()
	t.columnsCount = deserializer.Uint16()

	data := deserializer.RemainedData()
	for i := column.OrderId(0); i < t.columnsCount; i++ {
		t.columns[i], data = factory.DeserializeColumnDefinition(data)
	}
}

// BitmapLen returns number of bytes in tuple that constitute null bitmap
func (t *StandardTable) BitmapLen() uint16 {
	return uint16(math.Ceil(float64(t.columnsCount) / 8))
}

func (t *StandardTable) ColumnCount() uint16 {
	return t.columnsCount
}

func (t *StandardTable) ColumnName(columnId column.OrderId) string {
	return t.columnsNames[columnId]
}

func (t *StandardTable) ColumnId(name string) column.OrderId {
	return t.colNameIdMap[name]
}

func (t *StandardTable) ColumnsIds(names []string) []column.OrderId {
	colIds := make([]column.OrderId, 0, len(names))
	for i, name := range names {
		colIds[i] = t.colNameIdMap[name]
	}

	return colIds
}

func (t *StandardTable) ColumnParser(id column.OrderId) column.DataParser {
	return t.columns[id].DataParser()
}

func (t *StandardTable) ColumnParsers(ids []column.OrderId) []column.DataParser {
	parsers := make([]column.DataParser, 0, len(ids))
	for i, id := range ids {
		parsers[i] = t.columns[id].DataParser()
	}

	return parsers
}

func (t *StandardTable) ColumnSerializer(id column.OrderId) column.DataSerializer {
	return t.columns[id].DataSerializer()
}

func (t *StandardTable) ColumnSerializers(ids []column.OrderId) []column.DataSerializer {
	serializers := make([]column.DataSerializer, 0, len(ids))
	for i, id := range ids {
		serializers[i] = t.columns[id].DataSerializer()
	}

	return serializers
}

func (t *StandardTable) AllColumnSerializer() []column.DataSerializer {
	serializers := make([]column.DataSerializer, t.columnsCount)
	for i := 0; i < int(t.columnsCount); i++ {
		serializers[i] = t.columns[i].DataSerializer()
	}

	return serializers
}

func (t *StandardTable) GetColumn(index column.OrderId) column.ImmDefinition {
	return t.columns[index]
}

func (t *StandardTable) AddColumn(definition column.Definition) error {
	_, ok := t.colNameIdMap[definition.Name()]
	if ok {
		return errors.New("table already contains column with name:" + definition.Name())
	}
	t.columns = append(t.columns, definition)
	t.colNameIdMap[definition.Name()] = t.columnsCount
	t.columnsNames = append(t.columnsNames, definition.Name())
	t.columnsCount++

	return nil
}

func (t *StandardTable) RemoveColumn(name string) error {
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

	newColumnNames := make([]string, t.columnsCount-1)
	copy(newColumnNames[0:colToRemoveId], t.columnsNames[0:colToRemoveId])
	copy(newColumnNames[colToRemoveId:], t.columnsNames[colToRemoveId+1:])
	t.columnsNames = newColumnNames

	newColumns := make([]column.Definition, t.columnsCount-1)
	copy(newColumns[0:colToRemoveId], t.columns[0:colToRemoveId])
	copy(newColumns[colToRemoveId:], t.columns[colToRemoveId+1:])
	t.columns = newColumns
	t.columnsCount--

	return nil
}
