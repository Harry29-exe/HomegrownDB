package sql

type Table struct {
	columns map[string]Column
	name    string
	byteLen uint32
}

type Column struct {
	name    string
	fType   ColumnType
	offset  uint32
	byteLen uint16
}

type ColumnType = uint16

type Row = []byte

func (t *Table) GetOffset(columnName string) uint32 {
	return t.columns[columnName].offset
}
