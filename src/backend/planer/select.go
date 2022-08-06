package planer

import (
	"HomegrownDB/backend/analyser/anode"
	"HomegrownDB/backend/planer/plan"
)

// PlanSelect is mvc for now,
// it supports for now only following queries Select f1 f2 From t
// without Where, sorting, grouping and so on
func PlanSelect(node anode.Select) plan.Plan {
	if len(node.Tables.Tables) > 1 {
		panic("multiple tables select not supported yet!")
	}

	queryPlan := plan.Plan{
		Tables: []plan.Table{
			{
				TableId:     node.Tables.Tables[0].Table.TableId(),
				PlanTableId: 0,
			},
		},
	}

	fields := make([]plan.SelectedField, len(node.Fields.Fields))
	for i, field := range node.Fields.Fields {
		fields[i] = plan.SelectedField{
			Name: field.FieldName,
			FieldId: plan.FieldId{
				PlanTableId: 0,
				ColumnOrder: field.Column.GetColumnId(), //todo columnId what it is?
			},
		}
	}

	queryPlan.RootNode = plan.SelectFields{
		Fields: fields,
		Child: plan.SeqScan{
			Table:      node.Tables.Tables[0].Table,
			Conditions: nil,
		},
	}

	return queryPlan
}
