package executor

import (
	"HomegrownDB/backend/internal/executor/exenode"
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/backend/qrow"
)

func ExecPlan(plan plan.Plan) qrow.RowBuffer {
	root := plan.RootNode
	exeNode := delegateCreateExeNode(root)
	createChildrenNodes(exeNode, plan.RootNode)

	buffer := exeNode.Init(exenode.InitOptions{StoreAllRows: true})
	return buffer
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
