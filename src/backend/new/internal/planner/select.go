package planner

import "HomegrownDB/backend/new/internal/node"

var Select = _select{}

type _select struct{}

func (s _select) Plan(query node.Query, plan node.PlanedStmt) (node.Plan, error) {
	fromExpr := query.FromExpr
	switch len(fromExpr.FromList) {
	case 0:
		return s.planValStream(query, plan)
	case 1:
		return s.planSimpleSelect(query, plan)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func (s _select) planSimpleSelect(query node.Query, plan node.PlanedStmt) (node.Plan, error) {
	fromExpr := query.FromExpr
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

func (s _select) planValStream(query node.Query, stmt node.PlanedStmt) (node.Plan, error) {
	//todo implement me
	panic("Not implemented")
}
