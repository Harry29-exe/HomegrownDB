package schema

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/utils"
	"io/ioutil"
	"os"
)

type DBSchema struct {
	tables map[string]Table
}

var schema DBSchema

var dbObjectIdCounter = utils.NewLockCounter(uint64(0))
var lobIdCounter = utils.NewLockCounter(uint64(0))

func GetNextDbObjectId() uint64 {
	return dbObjectIdCounter.IncrementAndGet()
}

func GetNextLobId() uint64 {
	return lobIdCounter.IncrementAndGet()
}

func readDBSchema(dbHomePath string) {
	home := dbsystem.GetDBHomePath()
	tablesDir := home + "/" + dbsystem.TableDirname

	tables, err := ioutil.ReadDir(tablesDir)
	if err != nil {
		panic("Directory: " + dbsystem.TableDirname + " " +
			"does not exist in directory: " + dbsystem.GetDBHomePath())
	}

	schemaTables := map[string]Table{}

	for _, table := range tables {
		tableName := table.Name()
		data, err := os.ReadFile(tablesDir + "/" + tableName + "/" + dbsystem.TableInfoFilename)
		if err != nil {
			panic("File " + dbsystem.TableInfoFilename + " for table " + tableName + " does not exist.")
		}

		schemaTables[tableName] = *DeserializeTable(data)
	}
}
