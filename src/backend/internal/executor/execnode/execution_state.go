package execnode

import (
	"HomegrownDB/backend/internal/node"
)

type ExecPlanState[P node.Plan] struct {
	Plan  P
	Left  ExecNode
	Right ExecNode
}
