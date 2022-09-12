package stores

import "HomegrownDB/dbsystem/tx"

// TxStore todo finish
type TxStore interface {
	GetCtx(id tx.Id) tx.Ctx
	NewCtx() tx.Ctx
	FinishTx(id tx.Id)
}
