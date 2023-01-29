package table

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/hgtype/coltype"
	relation "HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/relation/table/column"
	"errors"
	"math"
)

var (
	_ Definition        = &StdTable{}
	_ relation.Relation = &StdTable{}
)

type StdTable struct {
	relation.BaseRelation
	columns  []column.WDef
	rColumns []column.Def
	name     string

	columnName_OrderMap map[string]column.Order
	columnsNames        []string
	columnsCount        uint16
}

func (t *StdTable) Hash() string {
	//TODO implement me
	panic("implement me")
}

func (t *StdTable) SetName(name string) {
	t.name = name
}

func (t *StdTable) Name() string {
	return t.name
}

func (t *StdTable) Serialize(serializer *bparse.Serializer) {
	relation.SerializeBaseRelation(&t.BaseRelation, serializer)
	serializer.MdString(t.name)
	serializer.Uint16(t.columnsCount)

	for _, col := range t.columns {
		col.Serialize(serializer)
	}
}

func (t *StdTable) Deserialize(deserializer *bparse.Deserializer) {
	t.BaseRelation = relation.DeserializeBaseRelation(deserializer)
	t.name = deserializer.MdString()
	t.columnsCount = deserializer.Uint16()

	t.columns = make([]column.WDef, t.columnsCount)
	t.rColumns = make([]column.Def, t.columnsCount)
	t.columnName_OrderMap = map[string]column.Order{}
	t.columnsNames = make([]string, t.columnsCount)
	for i := 0; i < int(t.columnsCount); i++ {
		col := column.Deserialize(deserializer)
		t.columns[col.Order()] = col
		t.rColumns[col.Order()] = col
		t.columnName_OrderMap[col.Name()] = col.Order()
		t.columnsNames[col.Order()] = col.Name()
	}
}

// BitmapLen returns number of bytes in tuple that constitute null bitmap
func (t *StdTable) BitmapLen() uint16 {
	return uint16(math.Ceil(float64(t.columnsCount) / 8))
}

func (t *StdTable) ColumnCount() uint16 {
	return t.columnsCount
}

func (t *StdTable) CTypePattern() []coltype.ColumnType {
	//todo implement me
	panic("Not implemented")
}

func (t *StdTable) ColumnName(columnId column.Order) string {
	return t.columnsNames[columnId]
}

func (t *StdTable) ColumnId(order column.Order) dbobj.OID {
	return t.columns[order].Id()
}

func (t *StdTable) ColumnOrder(name string) (order column.Order, ok bool) {
	order, ok = t.columnName_OrderMap[name]
	return
}

// todo array of ctypes?
func (t *StdTable) ColumnType(id column.Order) coltype.ColumnType {
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
func (t *StdTable) ColumnById(id dbobj.OID) column.Def {
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
