package page

import (
	"HomegrownDB/dbsystem/reldef"
	page "HomegrownDB/dbsystem/storage/page/internal"
	"HomegrownDB/dbsystem/storage/page/internal/data"
	"HomegrownDB/dbsystem/tx"
)

// -------------------------
//      Page
// -------------------------

type (
	TableRPage = data.RPage
	TablePage  = data.WPage
)

func InitNewTablePage(pageSlot []byte, table reldef.TableRDefinition, pageId page.Id) TablePage {
	return data.InitNewPage(pageSlot, table.OID(), pageId, data.PatternFromTable(table))
}

func AsTablePage(pageData []byte, pageId page.Id, table reldef.TableRDefinition) TablePage {
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

func PatternFromTable(table reldef.TableRDefinition) TuplePattern {
	return data.PatternFromTable(table)
}

func NewTuple(values [][]byte, pattern TuplePattern, tx tx.Tx) Tuple {
	return data.NewTuple(values, pattern, tx)
}

func NewTempTuple(values [][]byte, pattern TuplePattern) Tuple {
	return data.NewTuple(values, pattern, nil)
}
