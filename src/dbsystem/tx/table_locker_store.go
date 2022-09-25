package tx

import (
	"HomegrownDB/dbsystem/schema/table"
	"sync"
)

type TablesLocks interface {
	WLockTable(table.Id)
	RLockTable(table.Id)

	WUnlockTable(table.Id)
	RUnlockTable(table.Id)
}

type tableLockerStore struct {
	storeLock  *sync.RWMutex
	tableLocks map[table.Id]TableLocker
}
