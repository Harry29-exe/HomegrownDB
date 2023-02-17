package planner

import (
	"HomegrownDB/backend/internal/node"
	"log"
)

var CreateRelation = createRelation{}

type createRelation struct{}

func (createRelation) Plan(createTableNode node.CreateRelation, parentState State) (node.Plan, error) {
	var planNode node.Plan
	switch {
	case createTableNode.FutureTable != nil:
		planNode = node.NewCreateRelationPlan(createTableNode.FutureTable,
			parentState.NextPlanNodeId(),
			parentState.Query,
		)
	//case createTableNode.FutureIndex != nil:
	//	node.NewCreateRelationPlan(createTableNode.FutureIndex)
	default:
		log.Panicf("not implemented")
	}

	return planNode, nil
}
