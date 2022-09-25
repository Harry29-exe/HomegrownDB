package access

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/schema/table"
	"fmt"
	"sync"
)

var DBTableIOStore TableIOStore

func init() {
	pathToTableData := fmt.Sprintf("%s/%s", dbsystem.DBHomePath(), dbsystem.TableDirname)

	var err error
	DBTableIOStore, err = NewTableIOStore(pathToTableData, table.DBTableStore.AllTables())
	if err != nil {
		panic(fmt.Sprintf("could not initialize TableIOStore because:\n%s", err.Error()))
	}
}

type TableIOStore interface {
	GetTableIO(name string) (TableDataIO, error)
	TableIO(id table.Id) TableDataIO
	NewTableIO(definition table.Definition) error
	RemoveTableIO(definition table.Definition) error
}

func NewTableIOStore(tablesPath string, tables []table.Definition) (TableIOStore, error) {
	var err error

	ios := make([]TableDataIO, len(tables))
	for i, def := range tables {
		ios[i], err = SingleDiscTableDataIO(fmt.Sprintf("%s/%s", tablesPath, def.Name()))
		if err != nil {
			return nil, err
		}
	}

	return &tableIOStore{
		storeLock: &sync.RWMutex{},
		ios:       ios,
	}, nil
}

type tableIOStore struct {
	storeLock *sync.RWMutex
	ios       []TableDataIO
}

func (t *tableIOStore) GetTableIO(name string) (TableDataIO, error) {
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

func (t *tableIOStore) TableIO(id table.Id) TableDataIO {
	t.storeLock.RLock()
	defer t.storeLock.RUnlock()
	return t.ios[id]
}

func (t *tableIOStore) NewTableIO(def table.Definition) error {
	//TODO implement me
	panic("implement me")
}

func (t *tableIOStore) RemoveTableIO(def table.Definition) error {
	//TODO implement me
	panic("implement me")
}
