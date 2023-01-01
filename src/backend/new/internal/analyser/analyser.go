package analyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/analyser/seganalyser"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/dbsystem/relation/table"
)

func Analyse(stmt pnode.RawStmt, store table.Store) (node.Query, error) {
	ctx := anlsr.NewCtx(store)
	innerStmt := stmt.Stmt

	return seganalyser.Query.Analyse(innerStmt, ctx)
}
