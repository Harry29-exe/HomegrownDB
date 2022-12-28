package execnode

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/storage/dpage"
	"HomegrownDB/dbsystem/tx"
)

type ExecNode interface {
	Next() dpage.Tuple
	HasNext() bool
	Init(plan node.Plan) error
}

func Create(plan node.Plan) {

}

func newAbstractNode(txCtx tx.Ctx) abstractNode {
	return abstractNode{
		TxCtx: txCtx,
	}
}

type abstractNode struct {
	TxCtx tx.Ctx
}
