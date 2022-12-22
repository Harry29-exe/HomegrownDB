package table

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/relation"
	"errors"
	"math"
)

type StdTable struct {
	objectId relation.ID
	tableId  Id
	columns  []column.WDef
	rColumns []column.Def
	name     string

	nextColumnId        appsync.SyncCounter[column.Id]
	columnName_OrderMap map[string]column.Order
	columnsNames        []string
	columnsCount        uint16
}

func (t *StdTable) RelationID() relation.ID {
	return t.objectId
}

func (t *StdTable) Data() relation.Data {
	//TODO implement me
	panic("implement me")
}

func (t *StdTable) SetTableId(id Id) {
	t.tableId = id
}

func (t *StdTable) TableId() Id {
	return t.tableId
}

func (t *StdTable) Hash() string {
	//TODO implement me
	panic("implement me")
}

func (t *StdTable) SetRelationId(id relation.ID) {
	t.objectId = id
}

func (t *StdTable) SetName(name string) {
	t.name = name
}

func (t *StdTable) Name() string {
	return t.name
}

func (t *StdTable) Serialize() []byte {
	serializer := bparse.NewSerializer()

	serializer.Uint32(uint32(t.objectId))
	serializer.MdString(t.name)

	columnCount := uint16(len(t.columnName_OrderMap))
	serializer.Uint16(columnCount)
	for _, col := range t.columns {
		serializer.Append(col.Serialize())
	}

	return serializer.GetBytes()
}

func (t *StdTable) Deserialize(tableDef []byte) {
	deserializer := bparse.NewDeserializer(tableDef)
	t.objectId = relation.ID(deserializer.Uint32())
	t.name = deserializer.MdString()
	t.columnsCount = deserializer.Uint16()

	data := deserializer.RemainedData()
	for i := column.Order(0); i < t.columnsCount; i++ {
		t.columns[i], data = column.Serialize(data)
	}
}

// BitmapLen returns number of bytes in tuple that constitute null bitmap
func (t *StdTable) BitmapLen() uint16 {
	return uint16(math.Ceil(float64(t.columnsCount) / 8))
}

func (t *StdTable) ColumnCount() uint16 {
	return t.columnsCount
}

func (t *StdTable) CTypePattern() []hgtype.HGType {
	//todo implement me
	panic("Not implemented")
}

func (t *StdTable) ColumnName(columnId column.Order) string {
	return t.columnsNames[columnId]
}

func (t *StdTable) ColumnId(order column.Order) column.Id {
	return t.columns[order].Id()
}

func (t *StdTable) ColumnOrder(name string) (order column.Order, ok bool) {
	order, ok = t.columnName_OrderMap[name]
	return
}

// todo array of ctypes?
func (t *StdTable) ColumnType(id column.Order) hgtype.Wrapper {
	return t.columns[id].CType()
}

func (t *StdTable) ColumnByName(name string) (col column.Def, ok bool) {
	var id column.Order
	id, ok = t.columnName_OrderMap[name]
	if !ok {
		return nil, false
	}
	return t.columns[id], true
}

// ColumnById todo rewrite this: create columnId_Ordermap initialize it and use it
func (t *StdTable) ColumnById(id column.Id) column.Def {
	for _, def := range t.columns {
		if def.Id() == id {
			return def
		}
	}
	panic("no column with provided id")
}

func (t *StdTable) Column(index column.Order) column.Def {
	return t.columns[index]
}

func (t *StdTable) Columns() []column.Def {
	return t.rColumns
}

func (t *StdTable) AddColumn(definition column.WDef) error {
	_, ok := t.columnName_OrderMap[definition.Name()]
	if ok {
		return errors.New("table already contains column with name:" + definition.Name())
	}
	definition.SetId(t.nextColumnId.GetAndIncr())
	definition.SetOrder(t.columnsCount)

	t.columns = append(t.columns, definition)
	t.rColumns = append(t.rColumns, definition)
	t.columnName_OrderMap[definition.Name()] = t.columnsCount
	t.columnsNames = append(t.columnsNames, definition.Name())
	t.columnsCount++

	return nil
}

func (t *StdTable) RemoveColumn(name string) error {
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

func (t *StdTable) initInMemoryFields() {
	colCount := len(t.columns)
	t.columnsNames = make([]string, colCount)
	t.columnName_OrderMap = map[string]column.Order{}

	for i, col := range t.columns {
		t.columnsNames[i] = col.Name()
		t.columnName_OrderMap[col.Name()] = column.Order(i)
	}
	t.columnsCount = uint16(colCount)
}
