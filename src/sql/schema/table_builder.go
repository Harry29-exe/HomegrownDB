package schema

type TableBuilder struct {
	Table *Table
}

func NewTableBuilder(tableName string) *TableBuilder {
	return &TableBuilder{
		Table: &Table{
			objectId: 0,
			columns:  map[string]*Column{},
			colList:  make([]*Column, 8),
			name:     tableName,
			byteLen:  0,
		},
	}
}

func (tb *TableBuilder) Get() *Table {
	return tb.Table
}

func (tb *TableBuilder) AddColumn(column *Column) {
	tb.Table.colList = append(tb.Table.colList, column)
	tb.Table.columns[column.Name] = column
}

func (tb *TableBuilder) AddNewColumn(
	name string,
	columnType ColumnType,
	nullable bool,
	autoincrement bool) {

	column := &Column{
		Name:          name,
		Type:          columnType,
		Offset:        tb.calcNewColOffset(),
		Nullable:      nullable,
		Autoincrement: autoincrement,
	}

	tb.Table.colList = append(tb.Table.colList, column)
	tb.Table.columns[name] = column
}

func (tb *TableBuilder) calcNewColOffset() int32 {
	lastCol := tb.Table.colList[len(tb.Table.colList)]
	if lastCol.Offset < 0 || !lastCol.Type.IsFixedSize ||
		lastCol.Type.LobStatus != NEVER {
		return -1
	}

	return lastCol.Offset + int32(lastCol.Type.ByteLen)
}
