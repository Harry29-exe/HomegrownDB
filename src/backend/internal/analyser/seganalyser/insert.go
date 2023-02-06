package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anlsr"
	"HomegrownDB/backend/internal/analyser/seganalyser/typanlr"
	node2 "HomegrownDB/backend/internal/node"
	pnode2 "HomegrownDB/backend/internal/pnode"
	"HomegrownDB/backend/internal/sqlerr"
	"HomegrownDB/dbsystem/reldef/tabdef/column"
)

var Insert = insert{}

type insert struct{}

func (i insert) Analyse(stmt pnode2.InsertStmt, parentCtx anlsr.QueryCtx) (node2.Query, error) {
	query := node2.NewQuery(node2.CommandTypeInsert, stmt)
	currentCtx := parentCtx.CreateChildCtx(query)

	rte, err := RteRangeVar.Analyse(stmt.Relation, currentCtx)
	if err != nil {
		return nil, err
	}
	query.RTables = append(query.RTables, rte.Rte)
	query.ResultRel = rte.Rte.Id

	err = i.analyseInsertSource(stmt, currentCtx)
	if err != nil {
		return nil, err
	}

	err = i.analyseInputTargetList(stmt.Columns, currentCtx)
	if err != nil {
		return nil, err
	}
	err = i.extendWithDefaultEntries(currentCtx)
	if err != nil {
		return nil, err
	}
	return query, err
}

func (i insert) analyseInsertSource(insertStmt pnode2.InsertStmt, currentCtx anlsr.QueryCtx) error {
	srcNode := insertStmt.SrcNode
	if srcNode == nil {
		return sqlerr.NewIllegalPNodeErr(nil, "expected insert source")
	} else if srcNode.Tag() != pnode2.TagSelectStmt {
		return sqlerr.NewIllegalPNodeErr(srcNode, "not supported as insert source")
	}

	selectStmt := srcNode.(pnode2.SelectStmt)
	if selectStmt.Values != nil {
		valuesRTE, err := RteValues.Analyse(selectStmt.Values, currentCtx)
		if err != nil {
			return err
		}
		currentCtx.Query.AppendRTE(valuesRTE.Rte)
		currentCtx.Query.FromExpr = node2.NewFromExpr2(nil, valuesRTE.RteRefNode)
	} else {
		subquery, err := Select.Analyse(selectStmt, currentCtx)
		if err != nil {
			return err
		}
		subqueryRTE := node2.NewSubqueryRTE(currentCtx.RteIdCounter.Next(), subquery)
		currentCtx.Query.AppendRTE(subqueryRTE)
		currentCtx.Query.FromExpr = node2.NewFromExpr2(nil, subqueryRTE.CreateRef())
	}

	return nil
}

func (i insert) analyseInputTargetList(targets []pnode2.ResultTarget, currentCtx anlsr.QueryCtx) error {
	query := currentCtx.Query
	relation := i.getRelationRTE(query).Ref
	sourceRTE := i.getSourceRTE(query)

	columns := relation.Columns()
	targetList := make([]node2.TargetEntry, len(columns))
	for sourceColId, resultTarget := range targets {
		entry, err := TargetEntry.AnalyseForInsert(resultTarget, currentCtx)
		if err != nil {
			return err
		}

		destType := relation.Column(entry.AttribNo).CType()
		srcType := sourceRTE.ColTypes[sourceColId]
		if err = typanlr.IsAssignable(destType, srcType); err != nil {
			return err
		}

		entry.ExprToExec = node2.NewVar(sourceRTE.Id, column.Order(sourceColId), destType)
		targetList[entry.AttribNo] = entry
	}
	query.TargetList = targetList

	return nil
}

func (i insert) extendWithDefaultEntries(currentCtx anlsr.QueryCtx) error {
	query := currentCtx.Query
	relation := query.GetRTE(query.ResultRel).Ref
	columns := relation.Columns()

	for colOrder, colDef := range columns {
		if query.TargetList[colOrder] != nil {
			continue
		}
		query.TargetList[colOrder] = node2.NewTargetEntry(
			node2.NewConst(colDef.CType().Tag(), nil),
			node2.AttribNo(colOrder),
			colDef.Name(),
		)
	}

	return nil
}

func (i insert) getSourceRTE(insertQuery node2.Query) node2.RangeTableEntry {
	fromList := insertQuery.FromExpr.FromList
	if len(fromList) != 1 {
		panic("not supported from list len != 1 in insert query")
	} else if fromList[0].Tag() != node2.TagRteRef {
		panic("not supported from list node (only rte ref are supported)")
	}

	return insertQuery.GetRTE(fromList[0].(node2.RangeTableRef).Rte)
}

func (i insert) getRelationRTE(insertQuery node2.Query) node2.RangeTableEntry {
	return insertQuery.GetRTE(insertQuery.ResultRel)
}
