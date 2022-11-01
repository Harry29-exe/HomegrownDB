package tstructs

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/common/errors"
	"HomegrownDB/common/tests/tutils/testtable"
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/schema/table"
	"fmt"
	"sync"
	"testing"
)

func NewTestTableStoreWithInMemoryIO(t *testing.T, definitions ...testtable.TestTable) table.Store {
	definitionsMap := map[table.Id]testtable.TestTable{}
	tableIOs := map[table.Id]access.TableDataIO{}
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
		changeListeners: nil,
		tableIdCounter:  appsync.NewIdResolver(maxId+1, nil),
	}
}

type TestTablesStore struct {
	storeLock *sync.RWMutex

	nameTableMap map[string]table.Id
	definitions  map[table.Id]testtable.TestTable

	// store utils
	changeListeners []func()
	tableIdCounter  *appsync.IdResolver[table.Id]

	t *testing.T
}

func (i *TestTablesStore) GetTable(name string) (table.Definition, error) {
	i.storeLock.RLock()
	defer i.storeLock.RUnlock()

	id, ok := i.nameTableMap[name]
	if ok {
		return i.definitions[id], nil
	}
	return nil, fmt.Errorf("no table with name: %s", name)
}

func (i *TestTablesStore) Table(id table.Id) table.Definition {
	i.storeLock.RLock()
	defer i.storeLock.RUnlock()

	return i.definitions[id]
}

func (i *TestTablesStore) GetTestTable(name string) testtable.TestTable {
	i.storeLock.RLock()
	defer i.storeLock.RUnlock()

	id, ok := i.nameTableMap[name]
	if ok {
		return i.definitions[id]
	}
	panic(fmt.Sprintf("no table with name: %s", name))
}

func (i *TestTablesStore) TestTable(id table.Id) testtable.TestTable {
	i.storeLock.RLock()
	defer i.storeLock.RUnlock()

	testTable := i.definitions[id]
	if i != nil {
		return testTable
	}
	panic(fmt.Sprintf("no table with id: %d", id))
}

func (i *TestTablesStore) AllTables() []table.Definition {
	i.storeLock.RLock()
	defer i.storeLock.RUnlock()

	length := len(i.definitions)
	array := make([]table.Definition, length)
	for i, def := range i.definitions {
		array[i] = def
	}
	return array
}

func (i *TestTablesStore) AddTable(table table.WDefinition) error {
	i.storeLock.Lock()
	defer i.storeLock.Unlock()

	testTable := testtable.NewTestTable(table, i.t)
	id := i.tableIdCounter.NextId()
	testTable.SetTableId(id)
	i.nameTableMap[testTable.Name()] = id
	i.definitions[id] = testTable

	return nil
}

func (i *TestTablesStore) RemoveTable(id table.Id) error {
	i.storeLock.Lock()
	defer i.storeLock.Unlock()

	def, ok := i.definitions[id]
	if !ok {
		return errors.TableNotExist{}
	}

	delete(i.definitions, id)
	delete(i.nameTableMap, def.Name())

	return nil
}
