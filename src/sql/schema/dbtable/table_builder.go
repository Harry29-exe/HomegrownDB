package dbtable

import (
	"HomegrownDB/sql/schema/difinitions"
	"HomegrownDB/sql/schema/impl"
)

type TableBuilder struct {
	table   *impl.DbTable
	columns []*difinitions.Column
}

func NewTableBuilder(tableName string) *TableBuilder {
	return &TableBuilder{
		table: &impl.DbTable{
			objectId: 0,
			columns:  map[string]*difinitions.Column{},
			colList:  make([]*difinitions.Column, 8),
			name:     tableName,
			byteLen:  0,
		},
	}
}

func (tb *TableBuilder) Build() *impl.DbTable {
	//fixedSizeCols := make([]*schema.Column, 0, len(tb.columns))
	//fixedColsLen := uint32(0)
	//nonFixedSizeCols := make([]*schema.Column, 0, len(tb.columns))
	//
	//for _, column := range tb.columns {
	//	if column.Type.IsFixedSize {
	//		fixedSizeCols = append(fixedSizeCols, column)
	//		fixedColsLen += column.Type.
	//	} else {
	//		nonFixedSizeCols = append(nonFixedSizeCols, column)
	//	}
	//}
	//
	//
	//
	//table := tb.table
	//lastCol := table.columns[len(table.columns)-1]
	//if lastCol.Type {
	//
	//}
	//return tb.table
	return nil
}

func (tb *TableBuilder) AddColumn(column *difinitions.Column) {
	tb.columns = append(tb.columns, column)
}

func (tb *TableBuilder) AddNewColumn(
	name string,
	columnType difinitions.ColumnType,
	nullable bool,
	autoincrement bool) {

	column := &difinitions.Column{
		Name:          name,
		Type:          columnType,
		Offset:        tb.calcNewColOffset(),
		Nullable:      nullable,
		Autoincrement: autoincrement,
	}

	tb.columns = append(tb.columns, column)
}

func (tb *TableBuilder) calcNewColOffset() int32 {
	lastCol := tb.table.colList[len(tb.table.colList)]
	if lastCol.Offset < 0 || !lastCol.Type.IsFixedSize ||
		lastCol.Type.LobStatus != NEVER {
		return -1
	}

	return lastCol.Offset + int32(lastCol.Type.ByteLen)
}
