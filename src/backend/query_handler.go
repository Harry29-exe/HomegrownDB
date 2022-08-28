package backend

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/executor"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/planer"
	"HomegrownDB/backend/qrow"
)

func HandleQuery(query string) (qrow.RowBuffer, error) {
	parseTree, err := parser.Parse(query)
	if err != nil {
		return nil, err
	}
	analyserTree, err := analyser.Analyse(parseTree)
	if err != nil {
		return nil, err
	}
	plan, err := planer.Plan(analyserTree)
	if err != nil {
		return nil, err
	}

	return executor.ExecPlan(plan), nil
}
