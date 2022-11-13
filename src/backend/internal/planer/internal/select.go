package internal

import (
	"HomegrownDB/backend/internal/analyser/anode"
	plan "HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/backend/internal/shared/qctx"
)

var Select = _select{}

type _select struct{}

// Plan is mvc for now,
// it supports for now only following queries Select f1 f2 From t
// without Where, sorting, grouping and so on
func (s _select) Plan(node anode.Select, ctx qctx.QueryCtx) plan.Plan {
	if len(node.Tables) > 1 {
		panic("multiple tables select not supported yet!")
	}

	queryPlan := plan.NewPlan()

	fields := make([]qctx.QColumnId, len(node.Fields))
	for i, field := range node.Fields {
		fields[i] = qctx.QColumnId{
			QTableId: field.Table,
			ColOrder: field.Column,
		}
	}

	queryPlan.SetRootNode(plan.ReduceFields{
		Fields: fields,
		Child:  plan.NewSeqScan(ctx.QTCtx.GetTableByQTableId(node.Tables[0])),
	})

	return queryPlan
}
