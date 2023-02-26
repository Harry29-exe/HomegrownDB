package reldef

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/hgtype"
	"errors"
	"math"
)

var (
	_ TableDefinition = &Table{}
	_ Relation        = &Table{}
)

type Table struct {
	BaseRelation
	columns  []ColumnDefinition
	rColumns []ColumnRDefinition

	columnName_OrderMap map[string]Order
	columnsCount        uint16
}

func (t *Table) Hash() string {
	//TODO implement me
	panic("implement me")
}

func (t *Table) SetName(name string) {
	t.RelName = name
}

// BitmapLen returns number of bytes in tuple that constitute null bitmap
func (t *Table) BitmapLen() uint16 {
	return uint16(math.Ceil(float64(t.columnsCount) / 8))
}

func (t *Table) ColumnCount() uint16 {
	return t.columnsCount
}

// todo array of ctypes?
func (t *Table) ColumnType(id Order) hgtype.ColType {
	return t.columns[id].CType()
}

func (t *Table) ColumnByName(name string) (col ColumnRDefinition, ok bool) {
	var id Order
	id, ok = t.columnName_OrderMap[name]
	if !ok {
		return nil, false
	}
	return t.columns[id], true
}

// ColumnById todo rewrite this: create columnId_Ordermap initialize it and use it
func (t *Table) ColumnById(id hglib.OID) ColumnRDefinition {
	for _, def := range t.columns {
		if def.Id() == id {
			return def
		}
	}
	panic("no column with provided id")
}

func (t *Table) Column(index Order) ColumnRDefinition {
	return t.columns[index]
}

func (t *Table) Columns() []ColumnRDefinition {
	return t.rColumns
}

func (t *Table) AddNewColumn(definition ColumnDefinition) error {
	_, ok := t.columnName_OrderMap[definition.Name()]
	if ok {
		return errors.New("tabdef already contains column with name:" + definition.Name())
	}
	definition.SetOrder(t.columnsCount)

	t.columns = append(t.columns, definition)
	t.rColumns = append(t.rColumns, definition)
	t.columnName_OrderMap[definition.Name()] = t.columnsCount
	t.columnsCount++

	return nil
}

func (t *Table) AddColumn(definition ColumnDefinition) error {
	_, ok := t.columnName_OrderMap[definition.Name()]
	if ok {
		return errors.New("tabdef already contains column with name:" + definition.Name())
	}

	t.columns = append(t.columns, definition)
	t.rColumns = append(t.rColumns, definition)
	t.columnName_OrderMap[definition.Name()] = t.columnsCount
	t.columnsCount++

	return nil
}

func (t *Table) RemoveColumn(name string) error {
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

	copy(t.columns[colToRemoveId:], t.columns[colToRemoveId+1:])
	t.columns = t.columns[:len(t.columns)-1]
	t.rColumns = t.rColumns[:len(t.columns)-1]

	t.columnsCount--

	return nil
}

func (t *Table) initInMemoryFields() {
	colCount := len(t.columns)
	t.columnName_OrderMap = map[string]Order{}

	for i, col := range t.columns {
		t.columnName_OrderMap[col.Name()] = Order(i)
	}
	t.columnsCount = uint16(colCount)
}
