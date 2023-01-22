package execnode

import (
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/storage/page"
)

type ExecNode interface {
	Next() page.Tuple
	HasNext() bool
	Init(plan node.Plan) error
	Close() error
}

func Create(plan node.Plan) {

}
