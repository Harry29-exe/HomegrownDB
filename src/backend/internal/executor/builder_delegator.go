package executor

import (
	exenode2 "HomegrownDB/backend/internal/executor/exenode"
	"HomegrownDB/backend/internal/planer/plan"
)

func delegateCreateExeNode(node plan.Node) exenode2.ExeNode {
	builder, ok := exeNodeBuilders[node.Type()]
	if !ok {
		//todo implement me
		panic("Not implemented")
	}

	return builder.Build(node)
}

var exeNodeBuilders = map[plan.NodeType]exenode2.Builder{
	plan.SeqScanNode: exenode2.SeqScanBuilder,
}
