package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/analyser/seganalyser/typanlr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/dbsystem/hgtype"
	"fmt"
)

// -------------------------
//      RteValues
// -------------------------

type FutureType struct {
	Type hgtype.Tag
	Args any
}

var RteValues = rteValues{}

type rteValues struct{}

func (v rteValues) Analyse(pnodeValues [][]pnode.Node, query node.Query, ctx anlsr.Ctx) (RteResult, error) {
	values := make([][]node.Expr, len(pnodeValues))
	var err error

	firstRow, err := v.analyseFirstRow(pnodeValues[0], query, ctx)
	values[0] = firstRow
	for row := 1; row < len(values); row++ {
		values[row], err = v.analyseRow(pnodeValues[row], firstRow, query, ctx)
		if err != nil {
			return RteResult{}, err
		}
	}

	rte := node.NewValuesRTE(ctx.RteIdCounter.Next(), values)
	err = v.analyseTypes(rte)
	return NewSingleRteResult(rte), err
}

func (v rteValues) analyseFirstRow(row []pnode.Node, query node.Query, ctx anlsr.Ctx) ([]node.Expr, error) {
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

func (v rteValues) analyseRow(row []pnode.Node, firstRow []node.Expr, query node.Query, ctx anlsr.Ctx) ([]node.Expr, error) {
	resultRow := make([]node.Expr, len(row))
	for col := 0; col < len(row); col++ {
		aConst, err := ExprDelegator.DelegateAnalyse(row[col], query, ctx)
		if err != nil {
			return nil, err
		} else if aConst.Type() != firstRow[col].Type() {
			return nil, fmt.Errorf("incompatible types %s != %s", aConst.Type().ToStr(), firstRow[col].Type().ToStr())
		}

		resultRow[col] = aConst
	}
	return resultRow, nil
}

func (v rteValues) analyseTypes(rte node.RangeTableEntry) error {
	values := rte.ValuesList
	futureTypes := typanlr.CreateFutureTypes(values[0])
	for row := 1; row < len(values); row++ {
		err := futureTypes.UpdateTypes(values[row])
		if err != nil {
			return err
		}
	}

	colTypes := futureTypes.CreateTypes()
	rte.ColTypes = colTypes
	return nil
}
