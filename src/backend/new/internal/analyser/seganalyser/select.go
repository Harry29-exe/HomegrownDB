package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/dbsystem/schema/column"
)

var Select = _select{}

type _select struct{}

func (s _select) Analyse(stmt pnode.SelectStmt, ctx anlsr.Ctx) (node.Query, error) {
	query := node.NewQuery(node.CommandTypeSelect, stmt)
	query.Command = node.CommandTypeSelect

	var err error
	if stmt.Values != nil {
		err = s.analyseValuesSelect(stmt, query, ctx)
	} else {
		err = s.analyseStdSelect(stmt, query, ctx)
	}
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (s _select) analyseStdSelect(stmt pnode.SelectStmt, query node.Query, ctx anlsr.Ctx) error {
	err := FromDelegator.Analyse(stmt.From, query, ctx)
	if err != nil {
		return err
	}

	entries, err := TargetEntries.Analyse(stmt.Targets, query, ctx)
	if err != nil {
		return err
	}

	query.TargetList = entries

	return nil
}

func (s _select) analyseValuesSelect(stmt pnode.SelectStmt, query node.Query, ctx anlsr.Ctx) error {
	valuesRteResult, err := RteValues.Analyse(stmt.Values, query, ctx)
	rte, rteRef := valuesRteResult.Rte, valuesRteResult.RteRefNode
	if err != nil {
		return err
	}

	query.FromExpr = node.NewFromExpr2(nil, rteRef)
	query.RTables = append(query.RTables, rte)
	query.TargetList = make([]node.TargetEntry, len(rte.ColTypes))
	for col := 0; col < len(rte.ColTypes); col++ {
		colRef := node.NewVar(rte.Id, column.Order(col), rte.ColTypes[col])
		query.TargetList[col] = node.NewTargetEntry(colRef, uint16(col), "")
	}

	return nil
}

var SelectValidator = selectVldtr{}

type selectVldtr struct{}

func (s selectVldtr) Validate(query node.Query, ctx anlsr.Ctx) error {
	return nil
}
