package stores

import (
	"HomegrownDB/datastructs/appsync"
	"HomegrownDB/dbsystem/io"
	"HomegrownDB/dbsystem/schema/table"
)

var Tables = initTables()

func initTables() *TableStore {
	//todo implement reading files
	return &TableStore{}
}

type TableStore struct {
	nameTableMap    map[string]table.Id
	definitions     []table.WDefinition
	tableIOs        []io.TableDataIO
	changeListeners []func()
	tableIdCounter  appsync.SyncCounter[table.Id]
}

func NewTableStore(tables []table.WDefinition) *TableStore {
	tablesLen := len(tables)

	ids := map[table.Id]bool{}
	var maxId table.Id = 0
	for _, def := range tables {
		if def.TableId() > maxId {
			maxId = def.TableId()
		}
		ids[def.TableId()] = true
	}


	nameTableMap := make(map[string]table.Id, tablesLen)
	tablesIOs := make([]io.TableDataIO, tablesLen)

	for i, def := range tables {
		nameTableMap[def.Name()] = def.TableId()
		tablesIOs[]
	}

	return &TableStore{
		nameTableMap:    map[string]table.Id{},
		definitions:     nil,
		changeListeners: nil,
		tableIdCounter:  appsync.NewUint32SyncCounter(0),
	}
}

func NewEmptyTableStore() *TableStore {
	return &TableStore{
		nameTableMap:    map[string]table.Id{},
		definitions:     nil,
		changeListeners: nil,
		tableIdCounter:  appsync.NewUint32SyncCounter(0),
	}
}

func (t *TableStore) GetTable(name string) table.Definition {
	id, ok := t.nameTableMap[name]
	if ok {
		return t.definitions[id]
	}
	return nil
}

func (t *TableStore) GetTableIO(name string) io.TableDataIO {
	id, ok := t.nameTableMap[name]
	if ok {
		return t.tableIOs[id]
	}
}

func (t *TableStore) TableIO(id table.Id) io.TableDataIO {
	return t.tableIOs[id]
}

func (t *TableStore) Table(id table.Id) table.Definition {
	return t.definitions[id]
}

func (t *TableStore) AllTables() []table.Definition {
	length := len(t.definitions)
	allTablesList := make([]table.Definition, length)
	for i, def := range t.definitions {
		allTablesList[i] = def
	}

	return allTablesList
}

func (t *TableStore) AddTable(table table.WDefinition) error {
	//table.SetTableId()
	//todo implement me
	panic("Not implemented")
}

func (t *TableStore) RemoveTable(id table.Id) error {
	//todo implement me
	panic("Not implemented")
}

func (t *TableStore) RegisterChangeListener(fn func()) {
	t.changeListeners = append(t.changeListeners, fn)
}
