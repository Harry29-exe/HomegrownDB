package analyser

import (
	"HomegrownDB/backend/internal/analyser/analyse"
	"HomegrownDB/backend/internal/analyser/anlctx"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/pnode"
	"HomegrownDB/dbsystem/access/relation"
)

func Analyse(stmt pnode.RawStmt, store relation.AccessMngr) (node.Query, error) {
	ctx := anlctx.NewCtx(store)
	innerStmt := stmt.Stmt

	return analyse.Query.Analyse(innerStmt, ctx)
}
