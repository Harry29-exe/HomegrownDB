package page

import (
	"HomegrownDB/dbsystem/relation/dbobj"
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

func InitNewPage(pageSlot []byte, ownerId dbobj.OID, pageId page.Id, pattern TuplePattern) WPage {
	return data.InitNewPage(pageSlot, ownerId, pageId, pattern)
}

func InitNewTablePage(pageSlot []byte, table table.RDefinition, pageId page.Id) WPage {
	return data.InitNewPage(pageSlot, table.OID(), pageId, data.PatternFromTable(table))
}

func AsPage(pageData []byte, ownerId dbobj.OID, pageId Id, pattern TuplePattern) WPage {
	return data.AsPage(pageData, ownerId, pageId, pattern)
}

func AsTablePage(pageData []byte, pageId page.Id, table table.RDefinition) WPage {
	return data.AsPage(pageData, table.OID(), pageId, data.PatternFromTable(table))
}

// -------------------------
//      Tuple
// -------------------------

type (
	RTuple       = data.RTuple
	WTuple       = data.WTuple
	Tuple        = data.Tuple
	TuplePattern = data.TuplePattern
	PatternCol   = data.PatternCol
	ColumnInfo   = data.PatternCol
	TID          = data.TID
	TupleIndex   = data.TupleIndex
)

func NewTupleBuilder() data.TupleBuilder {
	return data.NewTupleBuilder()
}

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

var TupleDebugger = data.TupleDebugger
