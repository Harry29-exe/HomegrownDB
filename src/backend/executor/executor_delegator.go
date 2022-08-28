package executor

import (
	"HomegrownDB/backend/executor/exenode"
	"HomegrownDB/backend/planer/plan"
)

func delegateCreateExeNode(node plan.Node) exenode.ExeNode {
	builder, ok := exeNodeBuilders[node.Type()]
	if !ok {
		//todo implement me
		panic("Not implemented")
	}

	return builder.Build(node)
}

var exeNodeBuilders = map[plan.NodeType]exenode.Builder{
	plan.SeqScanNode: exenode.SeqScanBuilder,
}
