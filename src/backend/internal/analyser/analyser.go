package analyser

import (
	"HomegrownDB/backend/internal/analyser/anlsr"
	"HomegrownDB/backend/internal/analyser/seganalyser"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/pnode"
	"HomegrownDB/dbsystem/access/relation"
)

func Analyse(stmt pnode.RawStmt, store relation.AccessMngr) (node.Query, error) {
	ctx := anlsr.NewCtx(store)
	innerStmt := stmt.Stmt

	return seganalyser.Query.Analyse(innerStmt, ctx)
}
