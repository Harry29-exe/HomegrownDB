package stores

import (
	"HomegrownDB/datastructs/appsync"
	"HomegrownDB/dbsystem/io"
	"HomegrownDB/dbsystem/schema/table"
	"fmt"
	"sync"
)

type TableStore struct {
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

func (t *TableStore) GetTable(name string) table.Definition {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	id, ok := t.nameTableMap[name]
	if ok {
		return t.definitions[id]
	}
	return nil
}

func (t *TableStore) GetTableIO(name string) io.TableDataIO {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	id, ok := t.nameTableMap[name]
	if ok {
		return t.tableIOs[id]
	}
	panic(fmt.Sprintf("No table io with table name: %s", name))
}

func (t *TableStore) TableIO(id table.Id) io.TableDataIO {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()
	return t.tableIOs[id]
}

func (t *TableStore) Table(id table.Id) table.Definition {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()
	return t.definitions[id]
}

func (t *TableStore) AllTables() []table.Definition {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()

	length := len(t.definitions)
	allTablesList := make([]table.Definition, length)
	for i, def := range t.definitions {
		allTablesList[i] = def
	}

	return allTablesList
}

func (t *TableStore) AddTable(table table.WDefinition) error {
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

func (t *TableStore) RemoveTable(id table.Id) error {
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

func (t *TableStore) RegisterChangeListener(fn func()) {
	t.changeListeners = append(t.changeListeners, fn)
}
