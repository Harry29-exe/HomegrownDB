package planer

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/analyser/internal"
	"HomegrownDB/backend/internal/planer/internal/select"
	"HomegrownDB/backend/internal/planer/plan"
)

func Plan(tree analyser.Tree) (plan.Plan, error) {
	switch tree.RootType {
	case internal.Select:
		aNode, ok := tree.Root.(anode.Select)
		if !ok {
			//todo implement me
			panic("Not implemented")
		}
		return _select.Plan(aNode), nil
	default:
		//todo implement me
		panic("Not implemented")
	}
}
