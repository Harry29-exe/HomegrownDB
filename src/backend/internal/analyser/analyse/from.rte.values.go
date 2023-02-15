package analyse

import (
	"HomegrownDB/backend/internal/analyser/analyse/typanlr"
	"HomegrownDB/backend/internal/analyser/anlctx"
	node "HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/pnode"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"fmt"
)

// -------------------------
//      RteValues
// -------------------------

type FutureType struct {
	Type rawtype.Tag
	Args any
}

var RteValues = rteValues{}

type rteValues struct{}

func (v rteValues) Analyse(pnodeValues [][]pnode.Node, currentCtx anlctx.QueryCtx) (RteResult, error) {
	values := make([][]node.Expr, len(pnodeValues))
	var err error

	firstRow, err := v.analyseFirstRow(pnodeValues[0], currentCtx)
	values[0] = firstRow
	for row := 1; row < len(values); row++ {
		values[row], err = v.analyseRow(pnodeValues[row], firstRow, currentCtx)
		if err != nil {
			return RteResult{}, err
		}
	}

	rte := node.NewValuesRTE(currentCtx.RteIdCounter.Next(), values)
	err = v.analyseTypes(rte)
	return NewSingleRteResult(rte), err
}

func (v rteValues) analyseFirstRow(row []pnode.Node, currentCtx anlctx.QueryCtx) ([]node.Expr, error) {
	resultRow := make([]node.Expr, len(row))
	var err error
	for i := 0; i < len(row); i++ {
		resultRow[i], err = ExprDelegator.DelegateAnalyse(row[i], currentCtx)
		if err != nil {
			return nil, err
		}
	}
	return resultRow, nil
}

func (v rteValues) analyseRow(row []pnode.Node, firstRow []node.Expr, currentCtx anlctx.QueryCtx) ([]node.Expr, error) {
	resultRow := make([]node.Expr, len(row))
	for col := 0; col < len(row); col++ {
		aConst, err := ExprDelegator.DelegateAnalyse(row[col], currentCtx)
		if err != nil {
			return nil, err
		} else if aConst.TypeTag() != firstRow[col].TypeTag() {
			return nil, fmt.Errorf("incompatible types %s != %s", aConst.TypeTag().ToStr(), firstRow[col].TypeTag().ToStr())
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
