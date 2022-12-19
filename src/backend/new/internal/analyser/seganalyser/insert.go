package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
)

var Insert = insert{}

type insert struct{}

func (i insert) Analyse(stmt pnode.InsertStmt, ctx anlsr.Ctx) (node.Query, error) {
	query := node.NewQuery(node.CommandTypeInsert, stmt)

	rte, err := RteRangeVar.Analyse(stmt.Relation, ctx)
	if err != nil {
		return nil, err
	}
	query.RTables = append(query.RTables, rte.Rte)
	query.ResultRel = rte.Rte.Id

	entries, err := TargetEntries.Analyse(stmt.Columns, query, ctx)
	if err != nil {
		return nil, err
	}
	query.TargetList = entries

	err = FromDelegator.Analyse([]pnode.Node{stmt.SrcNode}, query, ctx)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (i insert) extendTargetEntries(query node.Query, ctx anlsr.Ctx) error {
	//tabDef :=
	//todo implement me
	panic("Not implemented")
}
