package planner

import (
	node2 "HomegrownDB/backend/internal/node"
	"errors"
	"fmt"
)

var Select = _select{}

type _select struct{}

func (s _select) Plan(query node2.Query, parentState State) (node2.Plan, error) {
	fromExpr := query.FromExpr
	if len(fromExpr.FromList) < 1 {
		return nil, errors.New("can not parse select query with empty from expr") // todo better err
	} else if len(fromExpr.FromList) > 1 {
		//todo implement me
		panic("Not implemented")
	}

	fromNode := fromExpr.FromList[0]
	switch fromNode.Tag() {
	case node2.TagRteRef:
		rteId := fromNode.(node2.RangeTableRef).Rte
		rte := query.GetRTE(rteId)
		if rte == nil {
			return nil, fmt.Errorf("no rte with id: %+v", rteId)
		}
		switch rte.Kind {
		case node2.RteRelation:
			return s.planSimpleSelect(query, parentState)
		case node2.RteValues:
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

func (s _select) planSimpleSelect(query node2.Query, parentState State) (node2.Plan, error) {
	fromExpr := query.FromExpr
	fromRoot := fromExpr.FromList[0]
	if fromRoot.Tag() != node2.TagRteRef {
		//todo implement me
		panic("Not implemented")
	}
	rteRef := fromRoot.(node2.RangeTableRef)

	seqScan := node2.NewSeqScan(parentState.NextPlanNodeId(), query)
	seqScan.RteId = rteRef.Rte
	seqScan.TargetList = query.TargetList

	parentState.AppendRTE(query.RTables...)
	return seqScan, nil
}

func (s _select) planValStream(query node2.Query, parentState State) (node2.Plan, error) {
	valuesRTE := query.RTables[0]
	valScan := node2.NewValueScan(parentState.NextPlanNodeId(), valuesRTE, query)
	valScan.RteId = valuesRTE.Id

	parentState.AppendRTE(valuesRTE)

	return valScan, nil
}
