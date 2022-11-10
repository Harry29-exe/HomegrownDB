package internal

import (
	"HomegrownDB/backend/internal/analyser/anode"
	plan "HomegrownDB/backend/internal/planer/plan"
)

// Plan is mvc for now,
// it supports for now only following queries Select f1 f2 From t
// without Where, sorting, grouping and so on
func Plan(node anode.Select) plan.Plan {
	if len(node.Tables.Tables) > 1 {
		panic("multiple tables select not supported yet!")
	}

	queryPlan := plan.NewPlan()

	fields := make([]plan.SelectedField, len(node.Fields.Fields))
	for i, field := range node.Fields.Fields {
		fields[i] = plan.SelectedField{
			Name: field.FieldAlias,
			FieldId: plan.FieldId{
				PlanTableId: 0,
				ColumnOrder: field.Column.Order(), //todo columnId what it is?
			},
		}
	}

	queryPlan.SetRootNode(plan.SelectFields{
		Fields: fields,
		Child: plan.SeqScan{
			Table:      node.Tables.Tables[0].Def.TableId(),
			Conditions: nil,
		},
	})

	return queryPlan
}
