package page

import (
	"HomegrownDB/dbsystem/relation/table"
	page "HomegrownDB/dbsystem/storage/page/internal"
	"HomegrownDB/dbsystem/storage/page/internal/data"
	"HomegrownDB/dbsystem/tx"
)

// -------------------------
//      Page
// -------------------------

type (
	RPage = data.RPage
	WPage = data.WPage
)

func InitNewPage(pattern TuplePattern, pageId page.Id, pageSlot []byte) WPage {
	return data.InitNewPage(pattern, pageId, pageSlot)
}

func AsPage(pageData []byte, pageId page.Id, pattern TuplePattern) WPage {
	return data.AsPage(pageData, pageId, pattern)
}

func AsTablePage(pageData []byte, pageId page.Id, table table.RDefinition) WPage {
	return data.AsPage(pageData, pageId, data.PatternFromTable(table))
}

// -------------------------
//      Tuple
// -------------------------

type (
	RTuple       = data.RTuple
	WTuple       = data.WTuple
	Tuple        = data.Tuple
	TuplePattern = data.TuplePattern
	ColumnInfo   = data.ColumnInfo
	TID          = data.TID
	TupleIndex   = data.TupleIndex
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
