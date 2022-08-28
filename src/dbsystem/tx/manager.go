package tx

import (
	"HomegrownDB/common/datastructs/appsync"
)

//todo add manager initalization from file

var Manager = &txManager{
	txIdCounter: appsync.NewInt32SyncCounter(0),
	txHistory:   map[Id]Status{},
}

type txManager struct {
	txIdCounter appsync.SyncCounter[int32]
	txHistory   map[Id]Status
	listeners   map[Id][]chan uint8
}

func (txm *txManager) New(level IsolationLevel) Tx {
	return Tx{
		Id:     Id(txm.txIdCounter.IncrementAndGet()),
		Level:  level,
		Status: InProgress,
	}
}

func (txm *txManager) TxStatus(id Id) Status {
	status, ok := txm.txHistory[id]
	if !ok {
		return NotStarted
	}

	return status
}

func (txm *txManager) Listen(txId Id) {

}
