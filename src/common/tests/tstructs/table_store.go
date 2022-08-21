package tstructs

import (
	"HomegrownDB/common/errors"
	"HomegrownDB/datastructs/appsync"
	"HomegrownDB/dbsystem/io"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/stores"
	"fmt"
	"sync"
)

func NewTestTableStore(definitions []TestTable, tablesIOs []io.TableDataIO) stores.Tables {
	definitionsMap := map[table.Id]TestTable{}
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
		tableIdCounter:  appsync.NewIdResolver(maxId+1, nil),
	}
}

func NewTestTableStoreWithInMemoryIO(definitions ...TestTable) stores.Tables {
	definitionsMap := map[table.Id]TestTable{}
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
		tableIdCounter:  appsync.NewIdResolver(maxId+1, nil),
	}
}

type TestTablesStore struct {
	storeLock *sync.RWMutex

	nameTableMap map[string]table.Id
	definitions  map[table.Id]TestTable
	tableIOs     map[table.Id]io.TableDataIO

	// store utils
	changeListeners []func()
	tableIdCounter  *appsync.IdResolver[table.Id]
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

func (t *TestTablesStore) GetTestTable(name string) TestTable {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	id, ok := t.nameTableMap[name]
	if ok {
		return t.definitions[id]
	}
	panic(fmt.Sprintf("no table with name: %s", name))
}

func (t *TestTablesStore) TestTable(id table.Id) TestTable {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	testTable := t.definitions[id]
	if t != nil {
		return testTable
	}
	panic(fmt.Sprintf("no table with id: %d", id))
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

	testTable := TestTable{table}
	id := t.tableIdCounter.NextId()
	testTable.SetTableId(id)
	t.nameTableMap[testTable.Name()] = id
	t.definitions[id] = testTable
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
