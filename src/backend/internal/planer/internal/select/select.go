package _select

import (
	"HomegrownDB/backend/internal/analyser/anode"
	plan2 "HomegrownDB/backend/internal/planer/plan"
)

// Plan is mvc for now,
// it supports for now only following queries Select f1 f2 From t
// without Where, sorting, grouping and so on
func Plan(node anode.Select) plan2.Plan {
	if len(node.Tables.Tables) > 1 {
		panic("multiple tables select not supported yet!")
	}

	queryPlan := plan2.Plan{
		Tables: []plan2.Table{
			{
				TableId:     node.Tables.Tables[0].Table.TableId(),
				PlanTableId: 0,
			},
		},
	}

	fields := make([]plan2.SelectedField, len(node.Fields.Fields))
	for i, field := range node.Fields.Fields {
		fields[i] = plan2.SelectedField{
			Name: field.FieldName,
			FieldId: plan2.FieldId{
				PlanTableId: 0,
				ColumnOrder: field.Column.GetColumnId(), //todo columnId what it is?
			},
		}
	}

	queryPlan.RootNode = plan2.SelectFields{
		Fields: fields,
		Child: plan2.SeqScan{
			Table:      node.Tables.Tables[0].Table.TableId(),
			Conditions: nil,
		},
	}

	return queryPlan
}
