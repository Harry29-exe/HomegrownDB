package node

import (
	"HomegrownDB/backend/planer/plan"
	"HomegrownDB/dbsystem/bdata"
)

type Node interface {
	HasNext() bool
	Next() bdata.Tuple

	NextBatch() []bdata.Tuple

	All() []bdata.Tuple
}

type PlanTableId = plan.TableId
