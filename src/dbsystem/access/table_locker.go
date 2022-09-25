package access

import (
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"sync"
)

type TableLocker interface {
	Id() table.Id
	RLock(ctx tx.Ctx)
	RUnlock(ctx tx.Ctx)
	WLock(ctx tx.Ctx)
	WUnlock(ctx tx.Ctx)
}

func NewTableLocker(tableId table.Id) TableLocker {
	return &tableLocker{
		id:                tableId,
		structLocker:      &sync.Mutex{},
		rLocksTxIds:       map[tx.Id]bool{},
		rLocksStartSignal: make(chan bool),
		rLocksInProgress:  0,
		waitingWLocks:     make([]tx.Id, 0, 5),
		wLockInProgress:   false,
		wLockInProgressId: 0,
		wLockStartSignal:  make(chan tx.Id),
	}
}

type tableLocker struct {
	id           table.Id
	structLocker sync.Locker

	rLocksTxIds       map[tx.Id]bool
	rLocksStartSignal chan bool
	rLocksInProgress  uint

	waitingWLocks     []tx.Id
	wLockInProgress   bool
	wLockInProgressId tx.Id
	wLockStartSignal  chan tx.Id
}

func (t *tableLocker) Id() table.Id {
	return t.id
}

func (t *tableLocker) RLock(ctx tx.Ctx) {
	t.structLocker.Lock()
	if len(t.waitingWLocks) == 0 && !t.wLockInProgress {
		t.rLocksInProgress++
		t.rLocksTxIds[ctx.Info.TxId()] = true
		t.structLocker.Unlock()
		return
	}
	t.structLocker.Unlock()

	<-t.rLocksStartSignal
	t.structLocker.Lock()
	t.rLocksInProgress++
	t.structLocker.Unlock()
	return
}

func (t *tableLocker) RUnlock(ctx tx.Ctx) {
	t.structLocker.Lock()
	t.rLocksInProgress--
	delete(t.rLocksTxIds, ctx.Info.TxId())

	t.tryUnlockingWLock()

	t.structLocker.Unlock()
}

func (t *tableLocker) WLock(ctx tx.Ctx) {
	t.structLocker.Lock()
	if t.rLocksInProgress == 0 && !t.wLockInProgress {
		t.wLockInProgress = true
		t.wLockInProgressId = ctx.Info.TxId()
		t.structLocker.Unlock()
		return
	}
	t.waitingWLocks = append(t.waitingWLocks, ctx.Info.TxId())
	t.structLocker.Unlock()

	for {
		txId := <-t.wLockStartSignal
		if txId == ctx.Info.TxId() {
			return
		}
	}
}

func (t *tableLocker) WUnlock(ctx tx.Ctx) {
	t.structLocker.Lock()
	t.wLockInProgress = false

	t.tryUnlockingWLock()
	t.tryUnlockingRLocks()
	t.structLocker.Unlock()
}

func (t *tableLocker) nextWLockTx() tx.Id {
	id := t.waitingWLocks[0]
	length := len(t.waitingWLocks)
	copy(t.waitingWLocks[:length-1], t.waitingWLocks[1:length])
	t.waitingWLocks = t.waitingWLocks[:length-1]

	return id
}

func (t *tableLocker) tryUnlockingWLock() {
	if t.rLocksInProgress == 0 && len(t.waitingWLocks) > 0 {

		wLockTxId := t.nextWLockTx()
		t.wLockInProgress = true
		t.wLockInProgressId = wLockTxId
		t.wLockStartSignal <- wLockTxId
	} else if t.rLocksInProgress == 1 && len(t.waitingWLocks) > 0 {

		wLockWaiting := t.waitingWLocks[0]
		if t.rLocksTxIds[wLockWaiting] {
			rLockTxId := t.nextWLockTx()
			t.wLockInProgress = true
			t.wLockInProgressId = rLockTxId
			t.wLockStartSignal <- rLockTxId
		}
	}
}

func (t *tableLocker) tryUnlockingRLocks() {
	if !t.wLockInProgress {
		t.rLocksStartSignal <- true
	}
}
