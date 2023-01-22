package page

import (
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/page/internal/data"
	"HomegrownDB/dbsystem/tx"
)

type (
	RTuple       = data.RTuple
	WTuple       = data.WTuple
	Tuple        = data.Tuple
	TuplePattern = data.TuplePattern
	ColumnInfo   = data.ColumnInfo
	TID          = data.TID
)

func NewPattern(columns []ColumnInfo) TuplePattern {
	return data.NewPattern(columns)
}

func PatternFromTable(table table.RDefinition) TuplePattern {
	return data.PatternFromTable(table)
}

func NewTuple(values [][]byte, pattern TuplePattern, tx tx.Tx) Tuple {
	return data.NewTuple(values, pattern, tx)
}

func NewTempTuple(values [][]byte, pattern TuplePattern) Tuple {
	return data.NewTuple(values, pattern, nil)
}
