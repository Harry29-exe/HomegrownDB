package backend

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/executor"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/planer"
	"HomegrownDB/backend/qrow"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
)

func HandleQuery(query string, txCtx *tx.Ctx) (qrow.RowBuffer, error) {
	parseTree, err := parser.Parse(query, txCtx)
	if err != nil {
		return nil, err
	}
	analyserTree, err := analyser.Analyse(parseTree, txCtx, table.DBTableStore)
	if err != nil {
		return nil, err
	}
	plan, err := planer.Plan(analyserTree)
	if err != nil {
		return nil, err
	}

	return executor.ExecPlan(plan), nil
}
