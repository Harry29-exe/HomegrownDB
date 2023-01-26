package backend

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/executor"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/planner"
	"HomegrownDB/dbsystem/hg/di"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

func Execute(query string, tx tx.Tx, container di.ExecutionContainer) ([]page.Tuple, error) {
	parseTree, err := parser.Parse(query)
	if err != nil {
		return nil, err
	}
	queryTree, err := analyser.Analyse(parseTree, container.TableStore)
	if err != nil {
		return nil, err
	}
	plan, err := planner.Plan(queryTree)
	if err != nil {
		return nil, err
	}

	return executor.Execute(plan, tx, container), nil
}
