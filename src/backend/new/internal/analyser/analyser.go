package analyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/analyser/seganalyser"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/dbsystem/schema/table"
)

func Analyse(stmt pnode.RawStmt, store table.Store) (node.Query, error) {
	ctx := anlsr.NewQCtx(store)

	innerStmt := stmt.Stmt
	switch innerStmt.Tag() {
	case pnode.TagSelectStmt:
		return seganalyser.Select.Analyse(innerStmt.(pnode.SelectStmt), ctx)
	case pnode.TagInsertStmt:
		return seganalyser.Insert.Analyse(innerStmt.(pnode.InsertStmt), ctx)
	default:
		//todo implement me
		panic("Not implemented")
	}
}
