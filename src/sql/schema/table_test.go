package schema

import "testing"

func TestTableSerialization(t *testing.T) {
	table := Table{
		objectId: 0,
		columns:  nil,
		name:     "",
		byteLen:  0,
	}
}

func createTestTable() *Table {
	table := Table{
		objectId: 44331,
		columns: map[string]Column{
			"column1": {
				Name:          "column1",
				Type:          *GetColumnType(Int2, nil),
				Offset:        0,
				Nullable:      false,
				Autoincrement: false,
			},
			"column2": {
				Name:   "column2",
				Type:   *GetColumnType(Int8, nil),
				Offset: 2,
			},
		},
		name:    "test_table",
		byteLen: 0,
	}
}
