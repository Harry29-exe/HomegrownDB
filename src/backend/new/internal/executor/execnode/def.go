package execnode

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/storage/tpage"
)

type ExecNode interface {
	Next() []tpage.Tuple
	Init(plan node.Plan) error
}

func Create(plan node.Plan) {

}
