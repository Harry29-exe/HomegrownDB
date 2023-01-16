package tx

import (
	"HomegrownDB/common/datastructs/appsync"
	"fmt"
	"sync"
)

type Manager interface {
	New(level IsolationLevel) Tx
	Get(id Id) Tx
	TxStatus(id Id) Status
	ListenFor(txId Id) Signal
}

type Signal = chan bool

func NewManager(startTxID Id) Manager {
	return &txManager{
		txIdCounter:   appsync.NewSyncCounter(startTxID),
		txHistory:     map[Id]Status{},
		listenersLock: &sync.Mutex{},
		listeners:     map[Id]Signal{},
	}
}

var _ Manager = &txManager{}

type txManager struct {
	txIdCounter appsync.SyncCounter[Id]
	txMap       map[Id]Tx
	txHistory   map[Id]Status

	listenersLock sync.Locker
	listeners     map[Id]Signal
}

func (txm *txManager) New(level IsolationLevel) Tx {
	return &StdTx{
		Id:       Id(txm.txIdCounter.IncrAndGet()),
		TxLevel:  level,
		TxStatus: InProgress,
	}
}

func (txm *txManager) Get(id Id) Tx {
	if tx, ok := txm.txMap[id]; ok {
		return tx
	} else {
		panic(fmt.Sprintf("not transaction with id: %d", id))
	}
}

func (txm *txManager) TxStatus(id Id) Status {
	status, ok := txm.txHistory[id]
	if !ok {
		return NotStarted
	}

	return status
}

func (txm *txManager) ListenFor(txId Id) Signal {
	txm.listenersLock.Lock()
	defer txm.listenersLock.Unlock()
	if signal, ok := txm.listeners[txId]; ok {
		return signal
	}
	signal := make(Signal)
	txm.listeners[txId] = signal
	return signal
}
