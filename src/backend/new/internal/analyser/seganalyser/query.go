package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
)

var Query = queryAnlr{}

type queryAnlr struct {
}

func (q queryAnlr) Analyse(stmt pnode.Node, ctx anlsr.Ctx) (node.Query, error) {
	switch stmt.Tag() {
	case pnode.TagSelectStmt:
		return Select.Analyse(stmt.(pnode.SelectStmt), ctx)
	case pnode.TagInsertStmt:
		return Insert.Analyse(stmt.(pnode.InsertStmt), ctx)
	default:
		//todo implement me
		panic("Not implemented")
	}
}
