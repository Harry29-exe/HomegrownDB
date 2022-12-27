package execnode

import (
	"HomegrownDB/backend/new/internal/executor/execnode/exexpr"
	"HomegrownDB/backend/new/internal/executor/exinfr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/storage/dpage"
	"HomegrownDB/dbsystem/tx"
)

var _ Builder = valuesScanBuilder{}

type valuesScanBuilder struct{}

func (v valuesScanBuilder) Create(plan node.Plan, ctx exinfr.ExCtx) ExecNode {
	vsPlan := plan.(node.ValueScan)
	outputPattern := exinfr.PatternFromTargetList(plan.Plan().TargetList)
	valuesRTE := ctx.GetRTE(vsPlan.RteId)
	innerPattern := exinfr.PattenFromRTE(valuesRTE)

	return &ValuesScan{
		TxCTX:         ctx.TxCtx,
		OutputPattern: outputPattern,
		InnerPattern:  innerPattern,
		Plan:          vsPlan,
		Values:        vsPlan.Values,
		rowCounter:    0,
	}
}

var _ ExecNode = &ValuesScan{}

type ValuesScan struct {
	TxCTX *tx.Ctx

	OutputPattern *dpage.TuplePattern
	InnerPattern  *dpage.TuplePattern
	Plan          node.ValueScan
	Values        [][]node.Expr
	rowCounter    uint
}

func (v *ValuesScan) Next() dpage.Tuple {
	innerTuple := v.tupleFromValues()
	nodeInput := exexpr.ExNodeInput{
		Plan:     v.Plan,
		Internal: innerTuple,
	}
	outputTupleValues := make([][]byte, len(v.Plan.TargetList))
	for i := 0; i < len(outputTupleValues); i++ {
		entry := v.Plan.TargetList[i]
		outputTupleValues[i] = exexpr.Execute(entry.ExprToExec, nodeInput)
	}

	return dpage.NewTuple(outputTupleValues, v.OutputPattern, v.TxCTX)
}

func (v *ValuesScan) Init(plan node.Plan) error {
	//TODO implement me
	panic("implement me")
}

func (v *ValuesScan) tupleFromValues() dpage.Tuple {
	tupleValues := make([][]byte, len(v.Values[0]))
	for col := 0; col < len(tupleValues); col++ {
		expr := v.Values[v.rowCounter][col]
		tupleValues[col] = exexpr.Execute(expr, exexpr.ExNodeInput{})
	}
	v.rowCounter++
	return dpage.NewTuple(tupleValues, v.InnerPattern, v.TxCTX)
}
