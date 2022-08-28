package planer

import (
	"HomegrownDB/backend/analyser"
	"HomegrownDB/backend/analyser/anode"
	"HomegrownDB/backend/planer/internal/select"
	"HomegrownDB/backend/planer/plan"
)

func Plan(tree analyser.Tree) (plan.Plan, error) {
	switch tree.RootType {
	case analyser.Select:
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
