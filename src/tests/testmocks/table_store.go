package testmocks

import (
	"HomegrownDB/datastructs/appsync"
	"HomegrownDB/dbsystem/io"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/stores"
	"HomegrownDB/errors"
	"fmt"
	"sync"
)

func NewTestTableStore(definitions []table.WDefinition, tablesIOs []io.TableDataIO) stores.Tables {
	definitionsMap := map[table.Id]table.WDefinition{}
	tableIOs := map[table.Id]io.TableDataIO{}
	nameTableMap := map[string]table.Id{}
	maxId := table.Id(0)
	for i, def := range definitions {
		id := def.TableId()
		if id > maxId {
			maxId = id
		}
		definitionsMap[id] = def
		tableIOs[id] = tablesIOs[i]
		nameTableMap[def.Name()] = id
	}

	return &TestTablesStore{
		storeLock:       &sync.RWMutex{},
		nameTableMap:    nameTableMap,
		definitions:     definitionsMap,
		tableIOs:        tableIOs,
		changeListeners: nil,
		tableIdCounter:  appsync.NewUint32SyncCounter(maxId + 1),
	}
}

func NewTestTableStoreWithInMemoryIO(definitions []table.WDefinition) stores.Tables {
	definitionsMap := map[table.Id]table.WDefinition{}
	tableIOs := map[table.Id]io.TableDataIO{}
	nameTableMap := map[string]table.Id{}
	maxId := table.Id(0)
	for _, def := range definitions {
		id := def.TableId()
		if id > maxId {
			maxId = id
		}
		definitionsMap[id] = def
		tableIOs[id] = NewInMemoryTableIO()
		nameTableMap[def.Name()] = id
	}

	return &TestTablesStore{
		storeLock:       &sync.RWMutex{},
		nameTableMap:    nameTableMap,
		definitions:     definitionsMap,
		tableIOs:        tableIOs,
		changeListeners: nil,
		tableIdCounter:  appsync.NewUint32SyncCounter(maxId + 1),
	}
}

type TestTablesStore struct {
	storeLock *sync.RWMutex

	nameTableMap map[string]table.Id
	definitions  map[table.Id]table.WDefinition
	tableIOs     map[table.Id]io.TableDataIO

	// store utils
	changeListeners []func()
	tableIdCounter  appsync.SyncCounter[uint32]
}

func (t *TestTablesStore) GetTable(name string) (table.Definition, error) {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	id, ok := t.nameTableMap[name]
	if ok {
		return t.definitions[id], nil
	}
	return nil, fmt.Errorf("no table with name: %s", name)
}

func (t *TestTablesStore) Table(id table.Id) table.Definition {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	return t.definitions[id]
}

func (t *TestTablesStore) AllTables() []table.Definition {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	length := len(t.definitions)
	array := make([]table.Definition, length)
	for i, def := range t.definitions {
		array[i] = def
	}
	return array
}

func (t *TestTablesStore) AddTable(table table.WDefinition) error {
	t.storeLock.Lock()
	defer t.storeLock.Unlock()

	id := t.tableIdCounter.GetAndIncrement()
	table.SetTableId(id)
	t.nameTableMap[table.Name()] = id
	t.definitions[id] = table
	tableIO := NewInMemoryTableIO()
	t.tableIOs[id] = tableIO

	return nil
}

func (t *TestTablesStore) RemoveTable(id table.Id) error {
	t.storeLock.Lock()
	defer t.storeLock.Unlock()

	def, ok := t.definitions[id]
	if !ok {
		return errors.TableNotExist{}
	}

	delete(t.definitions, id)
	delete(t.tableIOs, id)
	delete(t.nameTableMap, def.Name())

	return nil
}

func (t *TestTablesStore) GetTableIO(name string) (io.TableDataIO, error) {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	id, ok := t.nameTableMap[name]
	if ok {
		return t.tableIOs[id], nil
	}
	return nil, fmt.Errorf("no table with name: %s", name)
}

func (t *TestTablesStore) TableIO(id table.Id) io.TableDataIO {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	return t.tableIOs[id]
}
