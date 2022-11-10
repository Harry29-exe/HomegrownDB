package planer

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/planer/internal"

	"HomegrownDB/backend/internal/planer/plan"
)

func Plan(tree analyser.Tree) (plan.Plan, error) {
	switch tree.RootType {
	case analyser.RootTypeSelect:
		aNode, ok := tree.Root.(anode.Select)
		if !ok {
			//todo implement me
			panic("Not implemented")
		}
		return internal.Plan(aNode), nil
	default:
		//todo implement me
		panic("Not implemented")
	}
}
