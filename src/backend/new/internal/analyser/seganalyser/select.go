package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/dbsystem/relation/table/column"
)

var Select = _select{}

type _select struct{}

func (s _select) Analyse(stmt pnode.SelectStmt, parentCtx anlsr.QueryCtx) (node.Query, error) {
	query := node.NewQuery(node.CommandTypeSelect, stmt)
	query.Command = node.CommandTypeSelect
	currentCtx := parentCtx.CreateChildCtx(query)

	var err error
	if stmt.Values != nil {
		err = s.analyseValuesSelect(stmt, currentCtx)
	} else {
		err = s.analyseStdSelect(stmt, currentCtx)
	}
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (s _select) analyseStdSelect(stmt pnode.SelectStmt, currentCtx anlsr.QueryCtx) error {
	err := FromDelegator.Analyse(stmt.From, currentCtx)
	if err != nil {
		return err
	}

	err = s.analyseTargetEntries(stmt.Targets, currentCtx)
	if err != nil {
		return err
	}

	return nil
}

func (s _select) analyseValuesSelect(stmt pnode.SelectStmt, currentCtx anlsr.QueryCtx) error {
	valuesRteResult, err := RteValues.Analyse(stmt.Values, currentCtx)
	rte, rteRef := valuesRteResult.Rte, valuesRteResult.RteRefNode
	if err != nil {
		return err
	}

	query := currentCtx.Query
	query.FromExpr = node.NewFromExpr2(nil, rteRef)
	query.RTables = append(query.RTables, rte)
	query.TargetList = make([]node.TargetEntry, len(rte.ColTypes))
	for col := 0; col < len(rte.ColTypes); col++ {
		colRef := node.NewVar(rte.Id, column.Order(col), rte.ColTypes[col])
		query.TargetList[col] = node.NewTargetEntry(colRef, node.AttribNo(col), "")
	}

	return nil
}

// -------------------------
//      internal
// -------------------------

func (s _select) analyseTargetEntries(resTargets []pnode.ResultTarget, currentCtx anlsr.QueryCtx) error {
	entries := make([]node.TargetEntry, len(resTargets))

	for i, resTarget := range resTargets {
		entry, err := s.analyseTargetEntry(resTarget, currentCtx)
		if err != nil {
			return err
		}
		entries[i] = entry
	}

	currentCtx.Query.TargetList = entries
	return nil
}

func (s _select) analyseTargetEntry(
	resTarget pnode.ResultTarget,
	currentCtx anlsr.QueryCtx,
) (node.TargetEntry, error) {
	valExpr, err := ExprDelegator.DelegateAnalyse(resTarget.Val, currentCtx)
	if err != nil {
		return nil, err
	}

	attribNo := node.AttribNo(len(currentCtx.Query.TargetList))
	entry := node.NewTargetEntry(valExpr, attribNo, resTarget.Name)
	return entry, err
}

// -------------------------
//      SelectValidator
// -------------------------

var SelectValidator = selectVldtr{}

type selectVldtr struct{}

func (s selectVldtr) Validate(query node.Query, ctx anlsr.Ctx) error {
	return nil
}
