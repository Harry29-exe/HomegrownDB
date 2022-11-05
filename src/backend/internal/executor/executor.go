package executor

import (
	"HomegrownDB/backend/internal/executor/exenode"
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/dbsystem/access/dbbs"
)

func ExecPlan(plan plan.Plan) []dbbs.QRow {
	root := plan.RootNode()
	exeNode := exenode.Build(root, nil)

	return exeNode.All()
}
