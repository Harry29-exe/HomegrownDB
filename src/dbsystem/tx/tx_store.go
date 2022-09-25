package tx

import (
	"sync"
)

var DBTxStore TxStore //todo init

// TxStore todo finish
type TxStore interface {
	GetCtx(id Id) *Ctx
	NewCtx() *Ctx
	FinishTx(id Id)
}
type txStore struct {
	txStoreLock sync.RWMutex
	ctxMap      map[Id]*Ctx
	nextTxId    Id
}

func (t *txStore) GetCtx(id Id) *Ctx {
	t.txStoreLock.RLock()
	defer t.txStoreLock.Unlock()
	return t.ctxMap[id]
}

func (t *txStore) NewCtx() *Ctx {

	//ctx := tx.NewContext()
	//todo implement me
	panic("Not implemented")
}

func (t *txStore) FinishTx(id Id) {
	//TODO implement me
	panic("implement me")
}
