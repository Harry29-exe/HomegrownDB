package planer

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/planer/internal"
	"HomegrownDB/backend/internal/shared/qctx"

	"HomegrownDB/backend/internal/planer/plan"
)

func Plan(tree analyser.Tree, ctx qctx.QueryCtx) (plan.Plan, error) {
	switch tree.RootType {
	case analyser.RootTypeSelect:
		aNode, ok := tree.Root.(anode.Select)
		if !ok {
			//todo implement me
			panic("Not implemented")
		}
		return internal.Select.Plan(aNode, ctx), nil
	default:
		//todo implement me
		panic("Not implemented")
	}
}
