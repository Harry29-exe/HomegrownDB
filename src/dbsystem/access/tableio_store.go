package access

import (
	"HomegrownDB/dbsystem/schema/table"
	"sync"
)

type TablesIOs interface {
	GetTableIO(name string) (TableDataIO, error)
	TableIO(id table.Id) TableDataIO
	NewTableIO(definition table.Definition) error
	RemoveTableIO(definition table.Definition) error
}

type tablesIOs struct {
	storeLock *sync.RWMutex
	ios       []TableDataIO
}

func (t *tablesIOs) GetTableIO(name string) (TableDataIO, error) {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	//todo implement me
	panic("Not implemented")
	//id, ok := t.nameTableMap[name]
	//if ok {
	//	return t.ios[id], nil
	//}
	//return nil, fmt.Errorf("no table io with table name: %s", name)
}

func (t *tablesIOs) TableIO(id table.Id) TableDataIO {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()
	return t.ios[id]
}
