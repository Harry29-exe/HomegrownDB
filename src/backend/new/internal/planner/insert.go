package planner

import (
	"HomegrownDB/backend/new/internal/node"
	"errors"
)

var Insert = insert{}

type insert struct{}

func (i insert) Plan(query node.Query, plan node.PlanedStmt) (node.Plan, error) {
	insertPlan := node.NewModifyTable(plan.NextPlanNodeId(), node.ModifyTableInsert, query)
	insertPlan.TargetList = query.TargetList

	sourceRTE, err := i.retrieveSourceRTE(query)
	if err != nil {
		return nil, err
	}
	plan.AppendRteArr(query.RTables)

	srcPlan, err := delegate(sourceRTE.Subquery, plan)
	if err != nil {
		return nil, err
	}

	insertPlan.ResultRelations = []node.RteID{query.ResultRel}
	insertPlan.Left = srcPlan

	return insertPlan, nil
}

func (i insert) retrieveSourceRTE(query node.Query) (node.RangeTableEntry, error) {
	srcNode := query.FromExpr.FromList[0]
	if srcNode.Tag() != node.TagRteRef {
		return nil, errors.New("expected TagRteRef intead got: " + srcNode.Tag().ToString()) //todo better err
	}
	rteRef := srcNode.(node.RangeTableRef)
	rte, err := findRteWithId(rteRef.Rte, query.RTables)
	if err != nil {
		return nil, err
	} else if rte.Kind != node.RteSubQuery {
		return nil, errors.New("expected RteSubquery") //todo better err
	}
	return rte, nil
}

func findRteWithId(id node.RteID, rtes []node.RangeTableEntry) (node.RangeTableEntry, error) {
	for _, rte := range rtes {
		if rte.Id == id {
			return rte, nil
		}
	}
	return nil, errors.New("no rte with given id") //todo better err
}
