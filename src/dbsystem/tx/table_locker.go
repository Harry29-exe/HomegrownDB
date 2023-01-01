package tx

import (
	"HomegrownDB/dbsystem/relation/table"
	"sync"
)

type TableLocker interface {
	Id() table.Id
	RLock(ctx Ctx)
	RUnlock(ctx Ctx)
	WLock(ctx Ctx)
	WUnlock(ctx Ctx)
}

func NewTableLocker(tableId table.Id) TableLocker {
	return &tableLocker{
		id:                tableId,
		structLocker:      &sync.Mutex{},
		rLocksTxIds:       map[Id]bool{},
		rLocksStartSignal: make(chan bool),
		rLocksInProgress:  0,
		waitingWLocks:     make([]Id, 0, 5),
		wLockInProgress:   false,
		wLockInProgressId: 0,
		wLockStartSignal:  make(chan Id),
	}
}

type tableLocker struct {
	id           table.Id
	structLocker sync.Locker

	rLocksTxIds       map[Id]bool
	rLocksStartSignal chan bool
	rLocksInProgress  uint

	waitingWLocks     []Id
	wLockInProgress   bool
	wLockInProgressId Id
	wLockStartSignal  chan Id
}

func (t *tableLocker) Id() table.Id {
	return t.id
}

func (t *tableLocker) RLock(ctx Ctx) {
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

func (t *tableLocker) RUnlock(ctx Ctx) {
	t.structLocker.Lock()
	t.rLocksInProgress--
	delete(t.rLocksTxIds, ctx.Info.TxId())

	t.tryUnlockingWLock()

	t.structLocker.Unlock()
}

func (t *tableLocker) WLock(ctx Ctx) {
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

func (t *tableLocker) WUnlock(ctx Ctx) {
	t.structLocker.Lock()
	t.wLockInProgress = false

	t.tryUnlockingWLock()
	t.tryUnlockingRLocks()
	t.structLocker.Unlock()
}

func (t *tableLocker) nextWLockTx() Id {
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
