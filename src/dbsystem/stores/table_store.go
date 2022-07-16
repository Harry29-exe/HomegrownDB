package stores

import (
	"HomegrownDB/datastructs/appsync"
	"HomegrownDB/dbsystem/io"
	"HomegrownDB/dbsystem/schema/table"
	"sync"
)

var Tables = initTables()

func initTables() *TableStore {
	//todo implement reading files
	return &TableStore{}
}

func NewTableStore(tableDirectoryPath string, tables []table.WDefinition) (*TableStore, error) {
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

	return &TableStore{
		nameTableMap:    nameTableMap,
		definitions:     definitionsArray,
		tableIOs:        tablesIOs,
		changeListeners: nil,
		tableIdCounter:  appsync.NewIdResolver(maxId+1, missingIds),
	}, nil
}

func NewEmptyTableStore(tableDirectoryPath string) *TableStore {
	return &TableStore{
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
