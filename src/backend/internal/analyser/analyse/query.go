package analyse

import (
	"HomegrownDB/backend/internal/analyser/anlctx"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/pnode"
)

var Query = queryAnlr{}

type queryAnlr struct {
}

func (q queryAnlr) Analyse(stmt pnode.Node, ctx anlctx.Ctx) (node.Query, error) {
	rootCtx := anlctx.NewQueryCtx(nil, ctx)
	switch stmt.Tag() {
	case pnode.TagSelectStmt:
		return Select.Analyse(stmt.(pnode.SelectStmt), rootCtx)
	case pnode.TagInsertStmt:
		return Insert.Analyse(stmt.(pnode.InsertStmt), rootCtx)
	case pnode.TagCommandStmt:
		return CommandDelegator.Analyse(stmt.(pnode.CommandStmt), rootCtx)
	default:
		//todo implement me
		panic("Not implemented")
	}
}
