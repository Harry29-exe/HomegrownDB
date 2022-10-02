package executor

import (
	"HomegrownDB/backend/internal/executor/exenode"
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/dbsystem/dbbs"
)

func ExecPlan(plan plan.Plan) []dbbs.QRow {
	root := plan.RootNode
	exeNode := delegateCreateExeNode(root)
	createChildrenNodes(exeNode, plan.RootNode)

	return exeNode.All()
}

func createChildrenNodes(parent exenode.ExeNode, parentPlan plan.Node) {
	children := parentPlan.Children()
	if len(children) == 0 {
		return
	}

	nodeSrc := make([]exenode.ExeNode, len(children))
	for i := 0; i < len(children); i++ {
		nodeSrc[i] = delegateCreateExeNode(children[i])
		createChildrenNodes(nodeSrc[i], children[i])
	}
	parent.SetSource(nodeSrc)
}
