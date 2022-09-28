package table

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
	"errors"
	"math"
)

type StandardTable struct {
	objectId uint64
	tableId  Id
	columns  []column.WDef
	rColumns []column.Def
	name     string

	colNameIdMap map[string]column.OrderId
	columnsNames []string
	columnsCount uint16
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

func (t *StandardTable) OID() uint64 {
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
		t.columns[i], data = column.Serialize(data)
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

func (t *StandardTable) ColumnId(name string) (order column.OrderId, ok bool) {
	order, ok = t.colNameIdMap[name]
	return
}

func (t *StandardTable) ColumnsIds(names []string) []column.OrderId {
	colIds := make([]column.OrderId, 0, len(names))
	for i, name := range names {
		colIds[i] = t.colNameIdMap[name]
	}

	return colIds
}

func (t *StandardTable) ColumnType(id column.OrderId) ctype.CType {
	//todo implement me
	panic("Not implemented")
}

func (t *StandardTable) ColumnByName(name string) (col column.Def, ok bool) {
	var id column.OrderId
	id, ok = t.colNameIdMap[name]
	if !ok {
		return nil, false
	}
	return t.columns[id], true
}

func (t *StandardTable) Column(index column.OrderId) column.Def {
	return t.columns[index]
}

func (t *StandardTable) Columns() []column.Def {
	return t.rColumns
}

func (t *StandardTable) AddColumn(definition column.WDef) error {
	_, ok := t.colNameIdMap[definition.Name()]
	if ok {
		return errors.New("table already contains column with name:" + definition.Name())
	}
	t.columns = append(t.columns, definition)
	t.rColumns = append(t.rColumns, definition)
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

	delete(t.colNameIdMap, name)
	for colName, colId := range t.colNameIdMap {
		if colId > colToRemoveId {
			t.colNameIdMap[colName] = colId - 1
		}
	}

	copy(t.columnsNames[colToRemoveId:], t.columnsNames[colToRemoveId+1:])
	t.columnsNames = t.columnsNames[:len(t.columnsNames)-1]

	copy(t.columns[colToRemoveId:], t.columns[colToRemoveId+1:])
	t.columns = t.columns[:len(t.columns)-1]
	t.rColumns = t.rColumns[:len(t.columns)-1]

	t.columnsCount--

	return nil
}

func (t *StandardTable) initInMemoryFields() {
	colCount := len(t.columns)
	t.columnsNames = make([]string, colCount)
	t.colNameIdMap = map[string]column.OrderId{}

	for i, col := range t.columns {
		t.columnsNames[i] = col.Name()
		t.colNameIdMap[col.Name()] = column.OrderId(i)
	}
	t.columnsCount = uint16(colCount)
}
