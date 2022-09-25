package tstructs

import (
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/schema/table"
	"fmt"
	"sync"
)

type InMemTableIOStore struct {
	storeLock    *sync.RWMutex
	tableIOs     map[table.Id]access.TableDataIO
	nameTableMap map[string]table.Id
}

func NewInMemTableIO(definitions ...TestTable) access.TableIOStore {
	tableIOs := map[table.Id]access.TableDataIO{}
	nameTableMap := map[string]table.Id{}

	for _, def := range definitions {
		id := def.TableId()
		tableIOs[id] = NewInMemoryTableIO()
		nameTableMap[def.Name()] = id
	}

	return &InMemTableIOStore{
		storeLock:    &sync.RWMutex{},
		tableIOs:     tableIOs,
		nameTableMap: nameTableMap,
	}
}

func (i *InMemTableIOStore) GetTableIO(name string) (access.TableDataIO, error) {
	i.storeLock.RLock()
	defer i.storeLock.RUnlock()

	id, ok := i.nameTableMap[name]
	if ok {
		return i.tableIOs[id], nil
	}
	return nil, fmt.Errorf("no table with name: %s", name)
}

func (i *InMemTableIOStore) TableIO(id table.Id) access.TableDataIO {
	i.storeLock.RLock()
	defer i.storeLock.RUnlock()

	return i.tableIOs[id]
}

func (i *InMemTableIOStore) NewTableIO(definition table.Definition) error {
	//TODO implement me
	panic("implement me")
}

func (i *InMemTableIOStore) RemoveTableIO(definition table.Definition) error {
	//TODO implement me
	panic("implement me")
}
