package data

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"math"
)

type WithPattern interface {
	Pattern() TuplePattern
}

func PatternFromTable(def tabdef.TableRDefinition) TuplePattern {
	tableColumns := def.Columns()
	columns := make([]PatternCol, len(tableColumns))
	for c := 0; c < len(columns); c++ {
		columns[c] = PatternCol{
			Type: tableColumns[c].CType(),
			Name: tableColumns[c].Name(),
		}
	}

	return TuplePattern{
		Columns:   columns,
		BitmapLen: calcBitmapLen(len(columns)),
	}
}

func NewPattern(columns []PatternCol) TuplePattern {
	return TuplePattern{
		Columns:   columns,
		BitmapLen: calcBitmapLen(len(columns)),
	}
}

type TuplePattern struct {
	Columns   []PatternCol
	BitmapLen uint16
}

type PatternCol struct {
	Type hgtype.ColType
	Name string
}

func calcBitmapLen(columnCount int) uint16 {
	return uint16(math.Ceil(float64(columnCount) / 8))
}
