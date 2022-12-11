package planner

import "HomegrownDB/backend/new/internal/node"

func Plan(query node.Query) (node.PlanedStmt, error) {
	planedStmt := node.NewPlanedStmt(query.Command)

	var planTree node.Plan
	var err error
	switch query.Command {
	case node.CommandTypeSelect:
		planTree, err = Select.Plan(query, planedStmt)
	default:
		//todo implement me
		panic("Not implemented")
	}

	if err != nil {
		return planedStmt, err
	}
	planedStmt.PlanTree = planTree
	return planedStmt, nil
}
