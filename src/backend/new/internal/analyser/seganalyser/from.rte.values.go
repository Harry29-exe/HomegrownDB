package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/dbsystem/hgtype"
)

// -------------------------
//      RteValues
// -------------------------

type FutureType struct {
	Type hgtype.TypeTag
	Args any
}

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
