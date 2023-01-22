package dpage

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/relation/table"
	"math"
)

func NewPatternFromTable(def table.RDefinition) *TuplePattern {
	tableColumns := def.Columns()
	columns := make([]ColumnInfo, len(tableColumns))
	for c := 0; c < len(columns); c++ {
		columns[c] = ColumnInfo{
			CType: tableColumns[c].CType().Type,
			Name:  tableColumns[c].Name(),
		}
	}

	return &TuplePattern{
		Columns:   columns,
		BitmapLen: calcBitmapLen(len(columns)),
	}
}

func NewPattern(columns []ColumnInfo) *TuplePattern {
	return &TuplePattern{
		Columns:   columns,
		BitmapLen: calcBitmapLen(len(columns)),
	}
}

type TuplePattern struct {
	Columns   []ColumnInfo
	BitmapLen uint16
}

type ColumnInfo struct {
	CType hgtype.Type
	Name  string
}

func calcBitmapLen(columnCount int) uint16 {
	return uint16(math.Ceil(float64(columnCount) / 8))
}
