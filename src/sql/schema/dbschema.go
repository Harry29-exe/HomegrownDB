package schema

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/sql"
	"io/ioutil"
	"os"
)

type DBSchema struct {
	tables map[string]sql.Table
}

var schema DBSchema

func readDBSchema(dbHomePath string) {
	home := dbsystem.GetDBHomePath()
	tablesDir := home + "/" + dbsystem.TableDirname

	tables, err := ioutil.ReadDir(tablesDir)
	if err != nil {
		panic("Directory: " + dbsystem.TableDirname + " " +
			"does not exist in directory: " + dbsystem.GetDBHomePath())
	}

	schemaTables := map[string]sql.Table{}

	for _, table := range tables {
		tableName := table.Name()
		data, err := os.ReadFile(tablesDir + "/" + tableName + "/" + dbsystem.TableInfoFilename)
		if err != nil {
			panic("File " + dbsystem.TableInfoFilename + " for table " + tableName + " does not exist.")
		}

		schemaTables[tableName] = ReadTableSchema(data)
	}
}
