package schema

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/schema/table"
	"io/ioutil"
	"os"
)

var Tables = initTables()

type tables struct {
	definitions     []table.Definition
	changeListeners []func()
}

func (t *tables) Table(id table.Id) table.Definition {
	return t.definitions[id]
}

func (t *tables) AllTables() []table.Definition {
	length := len(t.definitions)
	allTablesList := make([]table.Definition, length)
	for i, def := range t.definitions {
		allTablesList[i] = def
	}

	return allTablesList
}

func (t *tables) RegisterChangeListener(fn func()) {
	t.changeListeners = append(t.changeListeners, fn)
}

func initTables() *tables {
	//todo implement reading files
	return &tables{}
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
