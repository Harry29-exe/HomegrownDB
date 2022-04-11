package schema

import (
	"HomegrownDB/datastructs"
	"HomegrownDB/dbsystem"
	"HomegrownDB/sql/schema/table"
	"io/ioutil"
	"os"
)

type dbSchema struct {
	tables map[string]table.Definition
}

var _ dbSchema

func (db *dbSchema) GetTable(name string) table.Definition {
	return db.tables[name]
}

var dbObjectIdCounter = datastructs.NewLockCounter(uint64(0))
var lobIdCounter = datastructs.NewLockCounter(uint64(0))

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
