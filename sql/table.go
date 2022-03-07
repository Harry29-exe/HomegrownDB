package sql

type Table struct {
	fields  []Field
	name    string
	byteLen uint32
	rows    uint64
}

type Field struct {
	name    string
	fType   FieldType
	byteLen uint16
}

type FieldType = uint16

type Row struct {
	data  []byte
	table *Table
}

func Serialize() {

}
