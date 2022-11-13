package internal

import (
	"HomegrownDB/backend/internal/analyser/anode"
	plan "HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/backend/internal/shared/qctx"
	"HomegrownDB/dbsystem/tx"
)

var Select = _select{}

type _select struct{}

// Plan is mvc for now,
// it supports for now only following queries Select f1 f2 From t
// without Where, sorting, grouping and so on
func (s _select) Plan(node anode.Select, tx tx.Ctx) plan.Plan {
	if len(node.Tables) > 1 {
		panic("multiple tables select not supported yet!")
	}

	queryPlan := plan.NewPlan()

	fields := make([]qctx.QFieldId, len(node.Fields))
	for i, field := range node.Fields {
		fields[i] = qctx.QFieldId{
			QTableId: field.Table,
			ColOrder: field.Column,
		}
	}

	queryPlan.SetRootNode(plan.ReduceFields{
		Fields: fields,
		Child: plan.SeqScan{
			Table:      tx.CurrentQuery.GetTableId(node.Tables[0]),
			Conditions: nil,
		},
	})

	return queryPlan
}
