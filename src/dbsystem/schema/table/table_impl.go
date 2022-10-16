package table

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/datastructs/appsync"
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

	nextColumnId        appsync.SyncCounter[column.Id]
	columnName_OrderMap map[string]column.Order
	columnsNames        []string
	columnsCount        uint16
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

	columnCount := uint16(len(t.columnName_OrderMap))
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
	for i := column.Order(0); i < t.columnsCount; i++ {
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

func (t *StandardTable) CTypePattern() []ctype.CType {
	//todo implement me
	panic("Not implemented")
}

func (t *StandardTable) ColumnName(columnId column.Order) string {
	return t.columnsNames[columnId]
}

func (t *StandardTable) ColumnId(order column.Order) column.Id {
	return t.columns[order].Id()
}

func (t *StandardTable) ColumnOrder(name string) (order column.Order, ok bool) {
	order, ok = t.columnName_OrderMap[name]
	return
}

// todo array of ctypes?
func (t *StandardTable) ColumnType(id column.Order) *ctype.CType {
	return t.columns[id].CType()
}

func (t *StandardTable) ColumnByName(name string) (col column.Def, ok bool) {
	var id column.Order
	id, ok = t.columnName_OrderMap[name]
	if !ok {
		return nil, false
	}
	return t.columns[id], true
}

// ColumnById todo rewrite this: create columnId_Ordermap initialize it and use it
func (t *StandardTable) ColumnById(id column.Id) column.Def {
	for _, def := range t.columns {
		if def.Id() == id {
			return def
		}
	}
	panic("no column with provided id")
}

func (t *StandardTable) Column(index column.Order) column.Def {
	return t.columns[index]
}

func (t *StandardTable) Columns() []column.Def {
	return t.rColumns
}

func (t *StandardTable) AddColumn(definition column.WDef) error {
	_, ok := t.columnName_OrderMap[definition.Name()]
	if ok {
		return errors.New("table already contains column with name:" + definition.Name())
	}
	definition.SetId(t.nextColumnId.GetAndIncrement())
	definition.SetOrder(t.columnsCount)

	t.columns = append(t.columns, definition)
	t.rColumns = append(t.rColumns, definition)
	t.columnName_OrderMap[definition.Name()] = t.columnsCount
	t.columnsNames = append(t.columnsNames, definition.Name())
	t.columnsCount++

	return nil
}

func (t *StandardTable) RemoveColumn(name string) error {
	colToRemoveId, ok := t.columnName_OrderMap[name]
	if !ok {
		return errors.New("column does not contain column with name: " + name)
	}

	delete(t.columnName_OrderMap, name)
	for colName, colId := range t.columnName_OrderMap {
		if colId > colToRemoveId {
			t.columnName_OrderMap[colName] = colId - 1
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
	t.columnName_OrderMap = map[string]column.Order{}

	for i, col := range t.columns {
		t.columnsNames[i] = col.Name()
		t.columnName_OrderMap[col.Name()] = column.Order(i)
	}
	t.columnsCount = uint16(colCount)
}
