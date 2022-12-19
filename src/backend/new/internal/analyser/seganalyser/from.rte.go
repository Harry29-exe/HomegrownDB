package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/dbsystem/hgtype"
)

// -------------------------
//      RteResult
// -------------------------

func NewSingleRteResult(rte node.RangeTableEntry) RteResult {
	return RteResult{
		Rte:        rte,
		RteRefNode: rte.CreateRef(),
	}
}

type RteResult struct {
	Rte        node.RangeTableEntry
	RteRefNode node.Node // RteRefNode is test like, node.RangeTableRef, Joins etc.
}

// -------------------------
//      RteRangeVar
// -------------------------

var RteRangeVar = rteRangeVar{}

type rteRangeVar struct{}

func (r rteRangeVar) Analyse(rangeVar pnode.RangeVar, ctx anlsr.Ctx) (RteResult, error) {
	def, err := ctx.GetTable(rangeVar.RelName)
	if err != nil {
		return RteResult{}, err
	}

	rte := node.NewRelationRTE(ctx.RteIdCounter.Next(), def)
	if rangeVar.Alias != "" {
		rte.Alias = node.NewAlias(rangeVar.Alias)
	}

	return NewSingleRteResult(rte), nil
}

// -------------------------
//      RteSubquery
// -------------------------

var RteSubquery = rteSelect{}

type rteSelect struct{}

func (rteSelect) Analyse(stmt pnode.SelectStmt, ctx anlsr.Ctx) (RteResult, error) {
	subquery, err := Select.Analyse(stmt, ctx)
	if err != nil {
		return RteResult{}, err
	}

	rte := node.NewSubqueryRTE(ctx.RteIdCounter.Next(), subquery)
	return NewSingleRteResult(rte), nil
}

// -------------------------
//      RteValues
// -------------------------

var RteValues = rteValues{}

type rteValues struct{}

func (v rteValues) Analyse(pnodeValues [][]pnode.Node, query node.Query, ctx anlsr.Ctx) (RteResult, error) {
	values := make([][]node.Expr, len(pnodeValues))
	var err error

	values[0], err = v.analyseRow(pnodeValues[0], query, ctx)

	for i := 1; i < len(values); i++ {
		values[i], err = v.analyseRow(pnodeValues[i], query, ctx)
		if err != nil {
			return RteResult{}, err
		}
	}

	rte := node.NewValuesRTE(ctx.RteIdCounter.Next(), values)
	return NewSingleRteResult(rte), nil
}

func (v rteValues) analyseRow(row []pnode.Node, query node.Query, ctx anlsr.Ctx) ([]node.Expr, error) {
	resultRow := make([]node.Expr, len(row))
	var err error
	for i := 0; i < len(row); i++ {
		resultRow[i], err = ExprDelegator.DelegateAnalyse(row[i], query, ctx)
		if err != nil {
			return nil, err
		}
	}
	return resultRow, nil
}

func (v rteValues) analyseFirstRow(
	row []pnode.Node,
	query node.Query,
	ctx anlsr.Ctx,
) (
	[]node.Expr,
	[]hgtype.Type,
) {
	resultRow := make([]node.Expr, len(row))
	types := make([]hgtype.Type)
	var err error

	for i := 0; i < len(row); i++ {
		resultRow[i], err = ExprDelegator.DelegateAnalyse(row[i], query, ctx)
		if err != nil {
			return nil, err
		}
	}
	return resultRow, nil
}
