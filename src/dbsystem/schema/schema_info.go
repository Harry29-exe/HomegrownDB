package schema

import (
	"HomegrownDB/datastructs/appsync"
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/schema/table"
	"io/ioutil"
	"os"
)

var Tables = initTables()

type tableSource struct {
	nameTableMap    map[string]table.Id
	definitions     []table.Definition
	changeListeners []func()
	tableIdCounter  appsync.SyncCounter[table.Id]
}

func NewTables() *tableSource {
	return &tableSource{
		nameTableMap:    map[string]table.Id{},
		definitions:     nil,
		changeListeners: nil,
		tableIdCounter:  appsync.NewUint32SyncCounter(0),
	}
}

func (t *tableSource) GetTable(name string) table.Definition {
	id, ok := t.nameTableMap[name]
	if ok {
		return t.definitions[id]
	}
	return nil
}

func (t *tableSource) Table(id table.Id) table.Definition {
	return t.definitions[id]
}

func (t *tableSource) AllTables() []table.Definition {
	length := len(t.definitions)
	allTablesList := make([]table.Definition, length)
	for i, def := range t.definitions {
		allTablesList[i] = def
	}

	return allTablesList
}

func (t *tableSource) AddTable(table table.WDefinition) error {
	//table.SetTableId()
	//todo implement me
	panic("Not implemented")
}

func (t *tableSource) RemoveTable(id table.Id) error {
	//todo implement me
	panic("Not implemented")
}

func (t *tableSource) RegisterChangeListener(fn func()) {
	t.changeListeners = append(t.changeListeners, fn)
}

func initTables() *tableSource {
	//todo implement reading files
	return &tableSource{}
}

func readDBSchema(dbHomePath string) {
	home := dbsystem.Props.DBHomePath()
	tablesDir := home + "/" + dbsystem.TableDirname

	tables, err := ioutil.ReadDir(tablesDir)
	if err != nil {
		panic("Directory: " + dbsystem.TableDirname + " " +
			"does not exist in directory: " + dbsystem.Props.DBHomePath())
	}

	schemaTables := map[string]table.Definition{}

	for _, tableInfo := range tables {
		tableName := tableInfo.Name()
		data, err := os.ReadFile(tablesDir + "/" + tableName + "/" + dbsystem.TableInfoFilename)
		if err != nil {
			panic("File " + dbsystem.TableInfoFilename + " for dbtable " + tableName + " does not exist.")
		}

		parsedTable := table.Deserializer.Deserialize(data)
		schemaTables[tableName] = parsedTable
	}

	//todo finish
}
