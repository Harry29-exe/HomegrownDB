package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
)

var Select = _select{}

type _select struct{}

func (s _select) Analyse(stmt pnode.SelectStmt, ctx anlsr.Ctx) (node.Query, error) {
	query := node.NewQuery(node.CommandTypeSelect, stmt)
	query.Command = node.CommandTypeSelect

	err := FromDelegator.Analyse(query, stmt.From, ctx)
	if err != nil {
		return nil, err
	}

	entries, err := TargetEntries.Analyse(stmt.Targets, query, ctx)
	if err != nil {
		return nil, err
	}

	query.TargetList = entries

	//todo implement me
	panic("Not implemented")
}

var SelectValidator = selectVldtr{}

type selectVldtr struct{}

func (s selectVldtr) Validate(query node.Query, ctx anlsr.Ctx) error {
	return nil
}
