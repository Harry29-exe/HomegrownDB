package planner

import (
	"HomegrownDB/backend/internal/node"
	"errors"
	"fmt"
)

var Select = _select{}

type _select struct{}

func (s _select) Plan(query node.Query, parentState State) (node.Plan, error) {
	fromExpr := query.FromExpr
	if len(fromExpr.FromList) < 1 {
		return nil, errors.New("can not parse select query with empty from expr") // todo better err
	} else if len(fromExpr.FromList) > 1 {
		//todo implement me
		panic("Not implemented")
	}

	fromNode := fromExpr.FromList[0]
	switch fromNode.Tag() {
	case node.TagRteRef:
		rteId := fromNode.(node.RangeTableRef).Rte
		rte := query.GetRTE(rteId)
		if rte == nil {
			return nil, fmt.Errorf("no rte with id: %+v", rteId)
		}
		switch rte.Kind {
		case node.RteRelation:
			return s.planSimpleSelect(query, parentState)
		case node.RteValues:
			return s.planValStream(query, parentState)
		default:
			//todo implement me
			panic("Not implemented")
		}
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func (s _select) planSimpleSelect(query node.Query, parentState State) (node.Plan, error) {
	fromExpr := query.FromExpr
	fromRoot := fromExpr.FromList[0]
	if fromRoot.Tag() != node.TagRteRef {
		//todo implement me
		panic("Not implemented")
	}
	rteRef := fromRoot.(node.RangeTableRef)

	seqScan := node.NewSeqScan(parentState.NextPlanNodeId(), query)
	seqScan.RteId = rteRef.Rte
	seqScan.TargetList = query.TargetList

	parentState.AppendRTE(query.RTables...)
	return seqScan, nil
}

func (s _select) planValStream(query node.Query, parentState State) (node.Plan, error) {
	valuesRTE := query.RTables[0]
	valScan := node.NewValueScan(parentState.NextPlanNodeId(), valuesRTE, query)
	valScan.RteId = valuesRTE.Id

	parentState.AppendRTE(valuesRTE)

	return valScan, nil
}
