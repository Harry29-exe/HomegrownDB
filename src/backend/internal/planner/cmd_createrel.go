package planner

import (
	"HomegrownDB/backend/internal/node"
)

var CreateRelation = createRelation{}

type createRelation struct{}

func (createRelation) Plan(createTableNode node.CreateRelation, parentState State) (node.Plan, error) {
	//todo implement me
	panic("Not implemented")
}
