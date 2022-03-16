package sql

type Table struct {
	cols    []Column
	name    string
	byteLen uint32
	rows    uint64
}

type Column struct {
	name    string
	fType   ColumnType
	byteLen uint16
}

type ColumnType = uint16

type Row struct {
	data  []byte
	table *Table
}

func Serialize() {

}
