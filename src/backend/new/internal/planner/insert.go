package planner

import (
	"HomegrownDB/backend/new/internal/node"
)

var Insert = insert{}

type insert struct{}

func (i insert) Plan(query node.Query, plan node.PlanedStmt) (node.Plan, error) {
	insertPlan := node.NewModifyTable(plan.NextPlanNodeId(), node.ModifyTableInsert, query)

	srcNode := query.FromExpr.FromList[0]
	srcPlan, err := delegate(srcNode.(node.Query), plan)
	if err != nil {
		return nil, err
	}

	insertPlan.ResultRelations = []node.RteID{query.ResultRel}
	insertPlan.Left = srcPlan

	return srcPlan, nil
}
