package execnode

import (
	"HomegrownDB/backend/internal/executor/execnode/exexpr"
	"HomegrownDB/backend/internal/executor/exinfr"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/storage/page"
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
		Tx:            ctx.Tx,
		OutputPattern: outputPattern,
		InnerPattern:  innerPattern,
		Plan:          vsPlan,
		Values:        vsPlan.Values,
		rowCounter:    0,
	}
}

var _ ExecNode = &ValuesScan{}

type ValuesScan struct {
	Tx tx.Tx

	OutputPattern page.TuplePattern
	InnerPattern  page.TuplePattern
	Plan          node.ValueScan
	Values        [][]node.Expr
	rowCounter    int
}

func (v *ValuesScan) Next() page.Tuple {
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

	return page.NewTuple(outputTupleValues, v.OutputPattern, v.Tx)
}

func (v *ValuesScan) HasNext() bool {
	return v.rowCounter < len(v.Values)
}

func (v *ValuesScan) Init(plan node.Plan) error {
	//TODO implement me
	panic("implement me")
}

func (v *ValuesScan) Shutdown() error {
	//TODO implement me
	panic("implement me")
}

func (v *ValuesScan) tupleFromValues() page.Tuple {
	tupleValues := make([][]byte, len(v.Values[0]))
	for col := 0; col < len(tupleValues); col++ {
		expr := v.Values[v.rowCounter][col]
		tupleValues[col] = exexpr.Execute(expr, exexpr.ExNodeInput{})
	}
	v.rowCounter++
	return page.NewTuple(tupleValues, v.InnerPattern, v.Tx)
}
