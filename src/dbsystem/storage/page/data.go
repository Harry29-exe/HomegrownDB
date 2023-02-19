package page

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/reldef/tabdef"
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

func InitNewPage(pageSlot []byte, ownerId hglib.OID, pageId page.Id, pattern TuplePattern) WPage {
	return data.InitNewPage(pageSlot, ownerId, pageId, pattern)
}

func InitNewTablePage(pageSlot []byte, table tabdef.RDefinition, pageId page.Id) WPage {
	return data.InitNewPage(pageSlot, table.OID(), pageId, data.PatternFromTable(table))
}

func AsPage(pageData []byte, ownerId hglib.OID, pageId Id, pattern TuplePattern) WPage {
	return data.AsPage(pageData, ownerId, pageId, pattern)
}

func AsTablePage(pageData []byte, pageId page.Id, table tabdef.RDefinition) WPage {
	return data.AsPage(pageData, table.OID(), pageId, data.PatternFromTable(table))
}

// -------------------------
//      Tuple
// -------------------------

type (
	RTuple       = data.RTuple
	WTuple       = data.WTuple
	Tuple        = data.Tuple
	TupleIndex   = data.TupleIndex
	TID          = data.TID
	TupleBuilder = data.TupleBuilder

	TuplePattern = data.TuplePattern
	PatternCol   = data.PatternCol
)

func NewTupleBuilder() data.TupleBuilder {
	return data.NewTupleBuilder()
}

func NewPattern(columns []PatternCol) TuplePattern {
	return data.NewPattern(columns)
}

func PatternFromTable(table tabdef.RDefinition) TuplePattern {
	return data.PatternFromTable(table)
}

func NewTuple(values [][]byte, pattern TuplePattern, tx tx.Tx) Tuple {
	return data.NewTuple(values, pattern, tx)
}

func NewTempTuple(values [][]byte, pattern TuplePattern) Tuple {
	return data.NewTuple(values, pattern, nil)
}
