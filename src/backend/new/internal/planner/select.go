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
	rteRef := fromRoot.(node.RangeTableRef)

	seqScan := node.NewSeqScan(plan.NextPlanNodeId(), query)
	seqScan.RteId = rteRef.Rte
	seqScan.TargetList = query.TargetList
	plan.Tables = append(plan.Tables, query.RTables...)

	return seqScan, nil
}
