package stores

import (
	"HomegrownDB/datastructs/appsync"
	"HomegrownDB/dbsystem/schema/table"
)

var Tables = initTables()

type TableStore struct {
	nameTableMap    map[string]table.Id
	definitions     []table.Definition
	changeListeners []func()
	tableIdCounter  appsync.SyncCounter[table.Id]
}

func NewTables() *TableStore {
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

func initTables() *TableStore {
	//todo implement reading files
	return &TableStore{}
}

//func readDBSchema(dbHomePath string) {
//	home := dbsystem.Props.DBHomePath()
//	tablesDir := home + "/" + dbsystem.TableDirname
//
//	tables, err := ioutil.ReadDir(tablesDir)
//	if err != nil {
//		panic("Directory: " + dbsystem.TableDirname + " " +
//			"does not exist in directory: " + dbsystem.Props.DBHomePath())
//	}
//
//	schemaTables := map[string]table.Definition{}
//
//	for _, tableInfo := range tables {
//		tableName := tableInfo.Name()
//		data, err := os.ReadFile(tablesDir + "/" + tableName + "/" + dbsystem.TableInfoFilename)
//		if err != nil {
//			panic("File " + dbsystem.TableInfoFilename + " for dbtable " + tableName + " does not exist.")
//		}
//
//		//parsedTable := table.Deserializer.Deserialize(data)
//		schemaTables[tableName] = parsedTable
//	}
//
//	//todo finish
//}
