package backend

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/executor"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/planer"
	"HomegrownDB/backend/internal/shared/qctx"
	"HomegrownDB/backend/internal/shared/query"
)

func HandleQuery(query string, ctx qctx.QueryCtx) ([]query.QRow, error) {
	parseTree, err := parser.Parse(query, ctx)
	if err != nil {
		return nil, err
	}
	analyserTree, err := analyser.Analyse(parseTree, ctx)
	if err != nil {
		return nil, err
	}
	plan, err := planer.Plan(analyserTree, ctx)
	if err != nil {
		return nil, err
	}

	return executor.ExecPlan(plan), nil
}
