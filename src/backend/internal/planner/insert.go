package planner

import (
	node2 "HomegrownDB/backend/internal/node"
	"errors"
)

var Insert = insert{}

type insert struct{}

func (i insert) Plan(query node2.Query, parentState State) (node2.Plan, error) {
	insertPlan := node2.NewModifyTable(parentState.NextPlanNodeId(), node2.ModifyTableInsert, query)
	insertPlan.ResultRelations = []node2.RteID{query.ResultRel}
	parentState.AppendRTE(query.RTables...)

	currentState := parentState.CreateChildState(query, insertPlan)
	err := i.handleInsertSource(currentState)
	if err != nil {
		return nil, err
	}

	return insertPlan, nil
}

func (i insert) handleInsertSource(currentState State) error {
	sourceRTE, err := i.retrieveSourceRTE(currentState.Query)
	if err != nil {
		return err
	}

	switch sourceRTE.Kind {
	case node2.RteValues:
		return i.handleRteValues(sourceRTE, currentState)
	case node2.RteSubQuery:
		panic("not implemented")
	default:
		panic("not supported")
	}
}

func (i insert) handleRteValues(sourceRTE node2.RangeTableEntry, currentState State) error {
	valueScan := node2.NewValueScan(
		currentState.NextPlanNodeId(),
		sourceRTE,
		currentState.Query,
	)
	valueScan.TargetList = currentState.Query.TargetList

	parentPlan := currentState.Plan.(node2.ModifyTable)
	parentPlan.Left = valueScan
	return nil
}

func (i insert) retrieveSourceRTE(query node2.Query) (node2.RangeTableEntry, error) {
	srcNode := query.FromExpr.FromList[0]
	if srcNode.Tag() != node2.TagRteRef {
		return nil, errors.New("expected TagRteRef intead got: " + srcNode.Tag().ToString()) //todo better err
	}
	rteRef := srcNode.(node2.RangeTableRef)
	rte := query.GetRTE(rteRef.Rte)
	if rte == nil {
		return nil, errors.New("insert has to have source rte")
	} else if rte.Kind != node2.RteSubQuery && rte.Kind != node2.RteValues {
		return nil, errors.New("expected RteSubquery or RteValues") //todo better err
	}
	return rte, nil
}
