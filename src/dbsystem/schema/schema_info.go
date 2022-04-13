package schema

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/schema/table"
	"io/ioutil"
	"os"
)

var Tables = initTables()

type tables struct {
	definitions table.Definition
}

func (t tables) Definition(name string) table.Definition {

}

func initTables() *tables {
	//todo implement reading files
	return &tables{}
}

func readDBSchema(dbHomePath string) {
	home := dbsystem.GetDBHomePath()
	tablesDir := home + "/" + dbsystem.TableDirname

	tables, err := ioutil.ReadDir(tablesDir)
	if err != nil {
		panic("Directory: " + dbsystem.TableDirname + " " +
			"does not exist in directory: " + dbsystem.GetDBHomePath())
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
