package seganalyser

import (
	anlsr2 "HomegrownDB/backend/internal/analyser/anlsr"
	node2 "HomegrownDB/backend/internal/node"
	pnode2 "HomegrownDB/backend/internal/pnode"
	"HomegrownDB/dbsystem/reldef/tabdef/column"
)

var Select = _select{}

type _select struct{}

func (s _select) Analyse(stmt pnode2.SelectStmt, parentCtx anlsr2.QueryCtx) (node2.Query, error) {
	query := node2.NewQuery(node2.CommandTypeSelect, stmt)
	query.Command = node2.CommandTypeSelect
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

func (s _select) analyseStdSelect(stmt pnode2.SelectStmt, currentCtx anlsr2.QueryCtx) error {
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

func (s _select) analyseValuesSelect(stmt pnode2.SelectStmt, currentCtx anlsr2.QueryCtx) error {
	valuesRteResult, err := RteValues.Analyse(stmt.Values, currentCtx)
	rte, rteRef := valuesRteResult.Rte, valuesRteResult.RteRefNode
	if err != nil {
		return err
	}

	query := currentCtx.Query
	query.FromExpr = node2.NewFromExpr2(nil, rteRef)
	query.RTables = append(query.RTables, rte)
	query.TargetList = make([]node2.TargetEntry, len(rte.ColTypes))
	for col := 0; col < len(rte.ColTypes); col++ {
		colRef := node2.NewVar(rte.Id, column.Order(col), rte.ColTypes[col])
		query.TargetList[col] = node2.NewTargetEntry(colRef, node2.AttribNo(col), "")
	}

	return nil
}

// -------------------------
//      internal
// -------------------------

func (s _select) analyseTargetEntries(resTargets []pnode2.ResultTarget, currentCtx anlsr2.QueryCtx) error {
	entries := make([]node2.TargetEntry, len(resTargets))

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
	resTarget pnode2.ResultTarget,
	currentCtx anlsr2.QueryCtx,
) (node2.TargetEntry, error) {
	valExpr, err := ExprDelegator.DelegateAnalyse(resTarget.Val, currentCtx)
	if err != nil {
		return nil, err
	}

	attribNo := node2.AttribNo(len(currentCtx.Query.TargetList))
	entry := node2.NewTargetEntry(valExpr, attribNo, resTarget.Name)
	return entry, err
}

// -------------------------
//      SelectValidator
// -------------------------

var SelectValidator = selectVldtr{}

type selectVldtr struct{}

func (s selectVldtr) Validate(query node2.Query, ctx anlsr2.Ctx) error {
	return nil
}
