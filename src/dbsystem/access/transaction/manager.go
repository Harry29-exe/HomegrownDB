package transaction

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/tx"
	"fmt"
	"sync"
)

type Manager interface {
	New(level tx.IsolationLevel) tx.Tx
	Get(id tx.Id) tx.Tx
	TxStatus(id tx.Id) tx.Status
	ListenFor(txId tx.Id) Signal
}

type Signal = chan bool

func NewManager(startTxID tx.Id) Manager {
	return &txManager{
		txIdCounter:   appsync.NewSyncCounter(startTxID),
		txHistory:     map[tx.Id]tx.Status{},
		listenersLock: &sync.Mutex{},
		listeners:     map[tx.Id]Signal{},
	}
}

var _ Manager = &txManager{}

type txManager struct {
	txIdCounter appsync.SyncCounter[tx.Id]
	txMap       map[tx.Id]tx.Tx
	txHistory   map[tx.Id]tx.Status

	listenersLock sync.Locker
	listeners     map[tx.Id]Signal
}

func (txm *txManager) New(level tx.IsolationLevel) tx.Tx {
	return &tx.StdTx{
		Id:       tx.Id(txm.txIdCounter.IncrAndGet()),
		TxLevel:  level,
		TxStatus: tx.InProgress,
	}
}

func (txm *txManager) Get(id tx.Id) tx.Tx {
	if tx, ok := txm.txMap[id]; ok {
		return tx
	} else {
		panic(fmt.Sprintf("not transaction with id: %d", id))
	}
}

func (txm *txManager) TxStatus(id tx.Id) tx.Status {
	status, ok := txm.txHistory[id]
	if !ok {
		return tx.NotStarted
	}

	return status
}

func (txm *txManager) ListenFor(txId tx.Id) Signal {
	txm.listenersLock.Lock()
	defer txm.listenersLock.Unlock()
	if signal, ok := txm.listeners[txId]; ok {
		return signal
	}
	signal := make(Signal)
	txm.listeners[txId] = signal
	return signal
}
