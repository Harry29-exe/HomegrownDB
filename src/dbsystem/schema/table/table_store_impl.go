package table

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/common/errors"
	"fmt"
	"sync"
)

type TablesStore struct {
	storeLock *sync.RWMutex

	// store data
	tableDirectoryPath string
	nameTableMap       map[string]Id
	definitions        []WDefinition

	// store utils
	changeListeners []func()
	tableIdCounter  *appsync.IdResolver[Id]
}

func NewTableStore(tableDirectoryPath string, tables []WDefinition) (*TablesStore, error) {
	maxId, missingIds := findMaxAndMissing(tables)

	nameTableMap := map[string]Id{}
	definitionsArray := make([]WDefinition, maxId)
	for _, def := range tables {
		id := def.TableId()
		nameTableMap[def.Name()] = id
		definitionsArray[id] = def
	}

	return &TablesStore{
		nameTableMap:    nameTableMap,
		definitions:     definitionsArray,
		changeListeners: nil,
		tableIdCounter:  appsync.NewIdResolver(maxId+1, missingIds),
	}, nil
}

func NewEmptyTableStore(tableDirectoryPath string) *TablesStore {
	return &TablesStore{
		storeLock: &sync.RWMutex{},

		tableDirectoryPath: tableDirectoryPath,
		nameTableMap:       map[string]Id{},
		definitions:        nil,

		changeListeners: nil,
		tableIdCounter:  appsync.NewIdResolver(Id(0), nil),
	}
}

func findMaxAndMissing(tables []WDefinition) (maxId Id, missing []Id) {
	maxId = Id(0)
	existingIds := map[Id]bool{}
	for _, def := range tables {
		if def.TableId() > maxId {
			maxId = def.TableId()
		}
		existingIds[def.TableId()] = true
	}

	for i := Id(0); i < maxId; i++ {
		if !existingIds[i] {
			missing = append(missing, i)
		}
	}

	return
}

func (t *TablesStore) GetTable(name string) (Definition, error) {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	id, ok := t.nameTableMap[name]
	if ok {
		return t.definitions[id], nil
	}
	return nil, errors.TableNotExist{TableName: name}
}

func (t *TablesStore) Table(id Id) Definition {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()
	return t.definitions[id]
}

func (t *TablesStore) AllTables() []Definition {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	length := len(t.definitions)
	allTablesList := make([]Definition, length)
	for i, def := range t.definitions {
		allTablesList[i] = def
	}

	return allTablesList
}

func (t *TablesStore) AddTable(table WDefinition) error {
	t.storeLock.Lock()
	defer t.storeLock.Unlock()

	id := t.tableIdCounter.NextId()
	table.SetTableId(id)
	if int(id) < len(t.definitions) {
		t.definitions[id] = table
	} else {
		t.definitions = append(t.definitions, table)
	}
	t.nameTableMap[table.Name()] = id

	return nil
}

func (t *TablesStore) RemoveTable(id Id) error {
	t.storeLock.Lock()
	defer t.storeLock.Unlock()

	if int(id) < len(t.definitions) {
		return fmt.Errorf("no table with id: %d", id)
	}
	tableDef := t.definitions[id]
	delete(t.nameTableMap, tableDef.Name())
	t.definitions[id] = nil
	t.tableIdCounter.AddId(id)

	return nil
}
