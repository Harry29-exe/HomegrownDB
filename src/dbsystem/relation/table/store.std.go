package table

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/relation"
	"fmt"
	"sync"
)

func NewTestTableStore(tables ...Definition) (Store, error) {
	return NewTableStore(tables)
}

func NewTableStore(tables []Definition) (Store, error) { //todo delete this error
	maxId, missingIds := findMaxAndMissing(tables)

	nameTableMap := map[string]Id{}
	definitionsArray := map[relation.ID]Definition{}
	for _, def := range tables {
		id := def.OID()
		nameTableMap[def.Name()] = id
		definitionsArray[id] = def
	}

	return &stdStore{
		nameTableMap:    nameTableMap,
		definitions:     definitionsArray,
		changeListeners: nil,
		tableIdCounter:  appsync.NewIdResolver(maxId+1, missingIds),
	}, nil
}

func NewEmptyTableStore() Store {
	return &stdStore{
		storeLock: &sync.RWMutex{},

		nameTableMap: map[string]Id{},
		definitions:  map[relation.ID]Definition{},

		changeListeners: nil,
		tableIdCounter:  appsync.NewIdResolver(Id(0), nil),
	}
}

func findMaxAndMissing(tables []Definition) (maxId Id, missing []Id) {
	maxId = Id(0)
	existingIds := map[Id]bool{}
	for _, def := range tables {
		if def.OID() > maxId {
			maxId = def.OID()
		}
		existingIds[def.OID()] = true
	}

	for i := Id(0); i < maxId; i++ {
		if !existingIds[i] {
			missing = append(missing, i)
		}
	}

	return
}

var _ Store = &stdStore{}

type stdStore struct {
	storeLock *sync.RWMutex

	// store data
	nameTableMap map[string]Id
	definitions  map[relation.ID]Definition

	// store utils
	changeListeners []func()
	tableIdCounter  *appsync.IdResolver[Id]
}

func (t *stdStore) FindTable(name string) Id {
	id, ok := t.nameTableMap[name]
	if !ok {
		return relation.InvalidRelId
	}
	return id
}

func (t *stdStore) AccessTable(id Id, lockMode TableLockMode) Definition {
	//todo add locking here
	return t.definitions[id]
}

func (t *stdStore) Table(id Id) RDefinition {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()
	return t.definitions[id]
}

func (t *stdStore) AddNewTable(table Definition) error {
	t.storeLock.Lock()
	defer t.storeLock.Unlock()

	id := t.tableIdCounter.NextId()
	table.SetOID(id)

	t.definitions[id] = table
	t.nameTableMap[table.Name()] = id

	return nil
}

func (t *stdStore) LoadTable(table Definition) error {
	t.storeLock.Lock()
	defer t.storeLock.Unlock()

	id := table.OID()
	t.definitions[id] = table
	t.nameTableMap[table.Name()] = id

	return nil
}

func (t *stdStore) RemoveTable(id Id) error {
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
