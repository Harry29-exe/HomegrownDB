package execnode

import (
	"HomegrownDB/backend/internal/executor/execnode/exexpr"
	exinfr2 "HomegrownDB/backend/internal/executor/exinfr"
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/storage/dpage"
	"HomegrownDB/dbsystem/tx"
)

var _ Builder = valuesScanBuilder{}

type valuesScanBuilder struct{}

func (v valuesScanBuilder) Create(plan node2.Plan, ctx exinfr2.ExCtx) ExecNode {
	vsPlan := plan.(node2.ValueScan)
	outputPattern := exinfr2.PatternFromTargetList(plan.Plan().TargetList)
	valuesRTE := ctx.GetRTE(vsPlan.RteId)
	innerPattern := exinfr2.PattenFromRTE(valuesRTE)

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

	OutputPattern *dpage.TuplePattern
	InnerPattern  *dpage.TuplePattern
	Plan          node2.ValueScan
	Values        [][]node2.Expr
	rowCounter    int
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

	return dpage.NewTuple(outputTupleValues, v.OutputPattern, v.Tx)
}

func (v *ValuesScan) HasNext() bool {
	return v.rowCounter < len(v.Values)
}

func (v *ValuesScan) Init(plan node2.Plan) error {
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
	return dpage.NewTuple(tupleValues, v.InnerPattern, v.Tx)
}
