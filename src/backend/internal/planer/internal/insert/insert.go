package insert

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/planer/plan"
)

func Plan(node anode.Insert) plan.Plan {
	p := plan.NewPlan()
	insertPlan := &plan.Insert{
		Table:   node.Table.Def,
		Columns: node.Columns,
	}
	p.AddTable(node.Table)

	if node.Expression != nil {
		panic("not supported yes (expression inside insert)")
	} else {
		insertPlan.Src = convertRowsIntoPlan(node.Rows, p)
	}

	return p
}

func convertRowsIntoPlan(rows []anode.InsertRow, p plan.Plan) plan.InsertNodeSrc {
	insertSrc := make([]plan.InsertRowSrc, len(rows))
	for i, row := range rows {
		insertSrc[i] = convertRow(row, p)
	}

	return plan.InsertNodeSrc{Rows: insertSrc}
}

func convertRow(row anode.InsertRow, p plan.Plan) plan.InsertRowSrc {
	insertSrc := plan.InsertRowSrc{Fields: make([]plan.InsertFieldSrc, len(row.Fields))}

	for i, field := range row.Fields {
		switch {
		case field.Expression != nil:
			panic("not supported yes (expression inside insert)")
		case field.Value != nil:
			insertSrc.Fields[i] = plan.InsertFieldSrc{
				Value: field.Value,
			}
		}
	}

	return insertSrc
}
