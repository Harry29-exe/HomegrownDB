package execnode

import (
	"HomegrownDB/backend/internal/executor/exinfr"
	"HomegrownDB/backend/internal/node"
	"fmt"
)

type Builder interface {
	Create(plan node.Plan, ctx exinfr.ExCtx) ExecNode
}

func CreateFromPlan(plan node.Plan, ctx exinfr.ExCtx) ExecNode {
	builder, ok := buildersMap[plan.Tag()]
	if !ok {
		panic(fmt.Sprintf("not supported tag: %s", plan.Tag().ToString()))
	}

	return builder.Create(plan, ctx)
}

var buildersMap map[node.Tag]Builder = map[node.Tag]Builder{
	node.TagValueScan:          valuesScanBuilder{},
	node.TagModifyTable:        modifyTableBuilder{},
	node.TagSeqScan:            seqScanBuilder{},
	node.TagCreateRelationPlan: createRelationBuilder{},
}
