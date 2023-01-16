package execnode

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/storage/dpage"
)

type ExecNode interface {
	Next() dpage.Tuple
	HasNext() bool
	Init(plan node.Plan) error
}

func Create(plan node.Plan) {

}
