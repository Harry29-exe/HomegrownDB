package planner

import (
	"HomegrownDB/backend/new/internal/node"
	"errors"
)

var Insert = insert{}

type insert struct{}

func (i insert) Plan(query node.Query, parentState State) (node.Plan, error) {
	insertPlan := node.NewModifyTable(parentState.NextPlanNodeId(), node.ModifyTableInsert, query)
	insertPlan.TargetList = query.TargetList
	insertPlan.ResultRelations = []node.RteID{query.ResultRel}
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
	case node.RteValues:
		return i.handleRteValues(sourceRTE, currentState)
	case node.RteSubQuery:
		panic("not implemented")
	default:
		panic("not supported")
	}
}

func (i insert) handleRteValues(sourceRTE node.RangeTableEntry, currentState State) error {
	valueScan := node.NewValueScan(
		currentState.NextPlanNodeId(),
		sourceRTE,
		currentState.Query,
	)

	parentPlan := currentState.Plan.(node.ModifyTable)
	parentPlan.Left = valueScan
	return nil
}

func (i insert) retrieveSourceRTE(query node.Query) (node.RangeTableEntry, error) {
	srcNode := query.FromExpr.FromList[0]
	if srcNode.Tag() != node.TagRteRef {
		return nil, errors.New("expected TagRteRef intead got: " + srcNode.Tag().ToString()) //todo better err
	}
	rteRef := srcNode.(node.RangeTableRef)
	rte := query.GetRTE(rteRef.Rte)
	if rte == nil {
		return nil, errors.New("insert has to have source rte")
	} else if rte.Kind != node.RteSubQuery && rte.Kind != node.RteValues {
		return nil, errors.New("expected RteSubquery or RteValues") //todo better err
	}
	return rte, nil
}
