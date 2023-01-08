package executor

import (
	"HomegrownDB/backend/new/internal/executor/execnode"
	"HomegrownDB/backend/new/internal/executor/exinfr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/dbsystem/storage/dpage"
	"HomegrownDB/dbsystem/tx"
)

func Execute(plan node.PlanedStmt, txCtx *tx.Ctx, dbStore hg.DBStore) []dpage.Tuple {
	ctx := exinfr.NewExCtx(plan, txCtx, dbStore)
	rootNode := execnode.CreateFromPlan(plan.PlanTree, ctx)

	tupleCache := make([]dpage.Tuple, 0, 100)
	for rootNode.HasNext() {
		tupleCache = append(tupleCache, rootNode.Next())
	}
	return tupleCache
}
