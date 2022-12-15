package planner

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/common/datastructs/appsync"
)

type PlanNodeIcCounter = appsync.SimpleSyncCounter[node.PlanNodeId]

func Plan(query node.Query) (node.PlanedStmt, error) {
	planedStmt := node.NewPlanedStmt(query.Command)

	planTree, err := delegate(query, planedStmt)

	if err != nil {
		return planedStmt, err
	}
	planedStmt.PlanTree = planTree
	return planedStmt, nil
}

func delegate(query node.Query, plan node.PlanedStmt) (node.Plan, error) {
	switch query.Command {
	case node.CommandTypeSelect:
		return Select.Plan(query, plan)
	case node.CommandTypeInsert:
		return Insert.Plan(query, plan)
	default:
		//todo implement me
		panic("Not implemented")
	}
}
