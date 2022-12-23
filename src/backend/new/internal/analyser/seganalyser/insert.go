package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"fmt"
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

func (i insert) extendSubquery(query node.Query, subquery node.Query, ctx anlsr.Ctx) error {
	err := i.validateAndPrepareSubquery(query, subquery)
	if err != nil {
		return err
	}

	tableDef := query.GetRTE(query.ResultRel).Ref
	queryTL := query.TargetList

	oldTlId := 0
	newInsertTL := make([]node.TargetEntry, len(queryTL))
	newSubqueryTL := make([]node.TargetEntry, len(queryTL))
	for col := uint16(0); col < uint16(len(queryTL)); col++ {
		if queryTL[oldTlId].AttribNo < col {
			columnDef := tableDef.Column(col)

			newInsertTL[col] = node.NewTargetEntry(
				node.NewConst(columnDef.CType().Tag, columnDef.DefaultValue()),
				col, columnDef.Name())
			newSubqueryTL[col] = node.NewTargetEntry(
				node.NewConst(columnDef.CType().Tag, columnDef.DefaultValue()),
				col, "")
		} else {
			oldTlId++
		}
	}
	query.TargetList = newInsertTL
	subquery.TargetList = newSubqueryTL
	return nil
}

func (i insert) validateAndPrepareSubquery(query, subquery node.Query) error {
	if len(query.TargetList) != len(subquery.TargetList) {
		return fmt.Errorf("expected %d columns but subquery has %d", len(query.TargetList), len(subquery.TargetList))
	}

	var insertEntry node.TargetEntry
	var srcEntry node.TargetEntry
	for col := 0; col < len(query.TargetList); col++ {
		insertEntry = query.TargetList[col]
		srcEntry = subquery.TargetList[col]
		if insertEntry.Type() != srcEntry.Type() {
			return fmt.Errorf("type %s can not be assigned to type %s",
				insertEntry.Type().ToStr(), srcEntry.Type().ToStr())
		}
		srcEntry.AttribNo = insertEntry.AttribNo
	}
	return nil
}
