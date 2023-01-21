package seganalyser

import (
	anlsr2 "HomegrownDB/backend/internal/analyser/anlsr"
	"HomegrownDB/backend/internal/node"
	pnode2 "HomegrownDB/backend/internal/pnode"
)

var Query = queryAnlr{}

type queryAnlr struct {
}

func (q queryAnlr) Analyse(stmt pnode2.Node, ctx anlsr2.Ctx) (node.Query, error) {
	rootCtx := anlsr2.NewQueryCtx(nil, ctx)
	switch stmt.Tag() {
	case pnode2.TagSelectStmt:
		return Select.Analyse(stmt.(pnode2.SelectStmt), rootCtx)
	case pnode2.TagInsertStmt:
		return Insert.Analyse(stmt.(pnode2.InsertStmt), rootCtx)
	default:
		//todo implement me
		panic("Not implemented")
	}
}
