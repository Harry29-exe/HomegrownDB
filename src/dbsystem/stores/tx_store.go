package stores

import (
	"HomegrownDB/dbsystem/tx"
	"sync"
)

var DBTxStore TxStore //todo init

// TxStore todo finish
type TxStore interface {
	GetCtx(id tx.Id) tx.Ctx
	NewCtx() tx.Ctx
	FinishTx(id tx.Id)
}
type txStore struct {
	txStoreLock sync.RWMutex
	ctxMap      map[tx.Id]tx.Ctx
	nextTxId    tx.Id
}

func (t *txStore) GetCtx(id tx.Id) tx.Ctx {
	t.txStoreLock.RLock()
	defer t.txStoreLock.Unlock()
	return t.ctxMap[id]
}

func (t *txStore) NewCtx() tx.Ctx {

	//ctx := tx.NewContext()
	//todo implement me
	panic("Not implemented")
}

func (t *txStore) FinishTx(id tx.Id) {
	//TODO implement me
	panic("implement me")
}
