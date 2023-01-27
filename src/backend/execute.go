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

func Execute(query string, tx tx.Tx, container di.ExecutionContainer) ([]page.RTuple, error) {
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

	return mapTuples(executor.Execute(plan, tx, container)), nil
}

func mapTuples(tuples []page.Tuple) []page.RTuple {
	rTuples := make([]page.RTuple, len(tuples))
	for i, tuple := range tuples {
		rTuples[i] = tuple
	}
	return rTuples
}
