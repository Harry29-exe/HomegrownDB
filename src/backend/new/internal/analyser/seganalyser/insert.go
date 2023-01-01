package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/analyser/seganalyser/typanlr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/backend/new/internal/sqlerr"
	"HomegrownDB/dbsystem/relation/table/column"
)

var Insert = insert{}

type insert struct{}

func (i insert) Analyse(stmt pnode.InsertStmt, parentCtx anlsr.QueryCtx) (node.Query, error) {
	query := node.NewQuery(node.CommandTypeInsert, stmt)
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

func (i insert) analyseInsertSource(insertStmt pnode.InsertStmt, currentCtx anlsr.QueryCtx) error {
	srcNode := insertStmt.SrcNode
	if srcNode == nil {
		return sqlerr.NewIllegalPNodeErr(nil, "expected insert source")
	} else if srcNode.Tag() != pnode.TagSelectStmt {
		return sqlerr.NewIllegalPNodeErr(srcNode, "not supported as insert source")
	}

	selectStmt := srcNode.(pnode.SelectStmt)
	if selectStmt.Values != nil {
		valuesRTE, err := RteValues.Analyse(selectStmt.Values, currentCtx)
		if err != nil {
			return err
		}
		currentCtx.Query.AppendRTE(valuesRTE.Rte)
		currentCtx.Query.FromExpr = node.NewFromExpr2(nil, valuesRTE.RteRefNode)
	} else {
		subquery, err := Select.Analyse(selectStmt, currentCtx)
		if err != nil {
			return err
		}
		subqueryRTE := node.NewSubqueryRTE(currentCtx.RteIdCounter.Next(), subquery)
		currentCtx.Query.AppendRTE(subqueryRTE)
		currentCtx.Query.FromExpr = node.NewFromExpr2(nil, subqueryRTE.CreateRef())
	}

	return nil
}

func (i insert) analyseInputTargetList(targets []pnode.ResultTarget, currentCtx anlsr.QueryCtx) error {
	query := currentCtx.Query
	relation := i.getRelationRTE(query).Ref
	sourceRTE := i.getSourceRTE(query)

	columns := relation.Columns()
	targetList := make([]node.TargetEntry, len(columns))
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

		entry.ExprToExec = node.NewVar(sourceRTE.Id, column.Order(sourceColId), destType)
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
		query.TargetList[colOrder] = node.NewTargetEntry(
			node.NewConst(colDef.CType().Tag, nil),
			node.AttribNo(colOrder),
			colDef.Name(),
		)
	}

	return nil
}

func (i insert) getSourceRTE(insertQuery node.Query) node.RangeTableEntry {
	fromList := insertQuery.FromExpr.FromList
	if len(fromList) != 1 {
		panic("not supported from list len != 1 in insert query")
	} else if fromList[0].Tag() != node.TagRteRef {
		panic("not supported from list node (only rte ref are supported)")
	}

	return insertQuery.GetRTE(fromList[0].(node.RangeTableRef).Rte)
}

func (i insert) getRelationRTE(insertQuery node.Query) node.RangeTableEntry {
	return insertQuery.GetRTE(insertQuery.ResultRel)
}
