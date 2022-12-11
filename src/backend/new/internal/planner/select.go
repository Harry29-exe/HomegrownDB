package planner

import "HomegrownDB/backend/new/internal/node"

var Select = _select{}

type _select struct{}

func (s _select) Plan(query node.Query, plan node.PlanedStmt) (node.Plan, error) {
	fromExpr := query.FromExpr
	if len(fromExpr.FromList) != 1 {
		//todo implement me
		panic("Not implemented")
	}

	fromRoot := fromExpr.FromList[0]
	if fromRoot.Tag() != node.TagRteRef {
		//todo implement me
		panic("Not implemented")
	}
	rteRef := fromRoot.(node.RangeTableEntry)

	seqScan := node.NewSeqScan(plan.NextPlanNodeId(), query)
	seqScan.RteId = rteRef.Id
	seqScan.TargetList = query.TargetList

	return seqScan, nil
}
