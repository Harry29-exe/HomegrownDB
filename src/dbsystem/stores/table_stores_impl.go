package stores

import (
	"HomegrownDB/datastructs/appsync"
	"HomegrownDB/dbsystem/io"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/errors"
	"fmt"
	"sync"
)

type TablesStore struct {
	storeLock *sync.RWMutex

	// store data
	tableDirectoryPath string
	nameTableMap       map[string]table.Id
	definitions        []table.WDefinition
	tableIOs           []io.TableDataIO

	// store utils
	changeListeners []func()
	tableIdCounter  *appsync.IdResolver[table.Id]
}

func NewTableStore(tableDirectoryPath string, tables []table.WDefinition) (*TablesStore, error) {
	maxId, missingIds := findMaxAndMissing(tables)

	nameTableMap := map[string]table.Id{}
	tablesIOs := make([]io.TableDataIO, maxId)
	definitionsArray := make([]table.WDefinition, maxId)
	for _, def := range tables {
		id := def.TableId()
		nameTableMap[def.Name()] = id
		definitionsArray[id] = def
		tableIO, err := io.SingleDiscTableDataIO(tableDirectoryPath + "/" + def.Name())
		if err != nil {
			return nil, err
		}
		tablesIOs[id] = tableIO
	}

	return &TablesStore{
		nameTableMap:    nameTableMap,
		definitions:     definitionsArray,
		tableIOs:        tablesIOs,
		changeListeners: nil,
		tableIdCounter:  appsync.NewIdResolver(maxId+1, missingIds),
	}, nil
}

func NewEmptyTableStore(tableDirectoryPath string) *TablesStore {
	return &TablesStore{
		storeLock: &sync.RWMutex{},

		tableDirectoryPath: tableDirectoryPath,
		nameTableMap:       map[string]table.Id{},
		definitions:        nil,
		tableIOs:           nil,

		changeListeners: nil,
		tableIdCounter:  appsync.NewIdResolver(table.Id(0), nil),
	}
}

func findMaxAndMissing(tables []table.WDefinition) (maxId table.Id, missing []table.Id) {
	maxId = table.Id(0)
	existingIds := map[table.Id]bool{}
	for _, def := range tables {
		if def.TableId() > maxId {
			maxId = def.TableId()
		}
		existingIds[def.TableId()] = true
	}

	for i := table.Id(0); i < maxId; i++ {
		if !existingIds[i] {
			missing = append(missing, i)
		}
	}

	return
}

func (t *TablesStore) GetTable(name string) (table.Definition, error) {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	id, ok := t.nameTableMap[name]
	if ok {
		return t.definitions[id], nil
	}
	return nil, errors.TableNotExist{TableName: name}
}

func (t *TablesStore) GetTableIO(name string) (io.TableDataIO, error) {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	id, ok := t.nameTableMap[name]
	if ok {
		return t.tableIOs[id], nil
	}
	return nil, fmt.Errorf("no table io with table name: %s", name)
}

func (t *TablesStore) TableIO(id table.Id) io.TableDataIO {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()
	return t.tableIOs[id]
}

func (t *TablesStore) Table(id table.Id) table.Definition {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()
	return t.definitions[id]
}

func (t *TablesStore) AllTables() []table.Definition {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	length := len(t.definitions)
	allTablesList := make([]table.Definition, length)
	for i, def := range t.definitions {
		allTablesList[i] = def
	}

	return allTablesList
}

func (t *TablesStore) AddTable(table table.WDefinition) error {
	tableIO, err := io.SingleDiscTableDataIO(t.tableDirectoryPath + "/" + table.Name())
	if err != nil {
		return err
	}

	t.storeLock.Lock()
	defer t.storeLock.Unlock()

	id := t.tableIdCounter.NextId()
	table.SetTableId(id)
	if int(id) < len(t.definitions) {
		t.definitions[id] = table
		t.tableIOs[id] = tableIO
	} else {
		t.definitions = append(t.definitions, table)
		t.tableIOs = append(t.tableIOs, tableIO)
	}
	t.nameTableMap[table.Name()] = id

	return nil
}

func (t *TablesStore) RemoveTable(id table.Id) error {
	t.storeLock.Lock()
	defer t.storeLock.Unlock()

	if int(id) < len(t.definitions) {
		return fmt.Errorf("no table with id: %d", id)
	}
	tableDef := t.definitions[id]
	delete(t.nameTableMap, tableDef.Name())
	t.definitions[id] = nil
	t.tableIOs[id] = nil
	t.tableIdCounter.AddId(id)

	return nil
}

func (t *TablesStore) RegisterChangeListener(fn func()) {
	t.changeListeners = append(t.changeListeners, fn)
}
