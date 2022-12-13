package executor

import (
	"HomegrownDB/backend/internal/executor/internal/exenode"
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/backend/internal/shared/query"
)

func ExecPlan(plan plan.Plan) []query.QRow {
	root := plan.RootNode()
	exeNode := exenode.Build(root, nil)

	return exeNode.All()
}
