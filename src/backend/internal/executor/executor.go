package executor

import (
	"HomegrownDB/backend/internal/executor/execnode"
	"HomegrownDB/backend/internal/executor/exinfr"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

func Execute(plan node.PlanedStmt, txCtx tx.Tx, dbStore hg.DBStore) []page.Tuple {
	ctx := exinfr.NewExCtx(plan, txCtx, dbStore)
	rootNode := execnode.CreateFromPlan(plan.PlanTree, ctx)

	tupleCache := make([]page.Tuple, 0, 100)
	for rootNode.HasNext() {
		tupleCache = append(tupleCache, rootNode.Next())
	}
	return tupleCache
}
