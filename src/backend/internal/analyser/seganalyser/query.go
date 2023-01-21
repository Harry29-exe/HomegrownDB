package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anlsr"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/pnode"
)

var Query = queryAnlr{}

type queryAnlr struct {
}

func (q queryAnlr) Analyse(stmt pnode.Node, ctx anlsr.Ctx) (node.Query, error) {
	rootCtx := anlsr.NewQueryCtx(nil, ctx)
	switch stmt.Tag() {
	case pnode.TagSelectStmt:
		return Select.Analyse(stmt.(pnode.SelectStmt), rootCtx)
	case pnode.TagInsertStmt:
		return Insert.Analyse(stmt.(pnode.InsertStmt), rootCtx)
	default:
		//todo implement me
		panic("Not implemented")
	}
}
