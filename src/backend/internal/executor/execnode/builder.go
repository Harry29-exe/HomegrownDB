package execnode

import (
	"HomegrownDB/backend/internal/executor/exinfr"
	node2 "HomegrownDB/backend/internal/node"
	"fmt"
)

type Builder interface {
	Create(plan node2.Plan, ctx exinfr.ExCtx) ExecNode
}

func CreateFromPlan(plan node2.Plan, ctx exinfr.ExCtx) ExecNode {
	builder, ok := buildersMap[plan.Tag()]
	if !ok {
		panic(fmt.Sprintf("not supported tag: %s", plan.Tag().ToString()))
	}

	return builder.Create(plan, ctx)
}

var buildersMap map[node2.Tag]Builder = map[node2.Tag]Builder{
	node2.TagValueScan:   valuesScanBuilder{},
	node2.TagModifyTable: modifyTableBuilder{},
}
