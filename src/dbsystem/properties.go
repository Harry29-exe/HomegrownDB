package dbsystem

import (
	"log"
	"os"
)

const dbHomeVarName = "HOMEGROWN_DB_HOME"
const TableDirname = "tables"
const TableInfoFilename = "info"

var dbHomePath = ""

var Props = &dbProperties{}

type dbProperties struct{}

func (dbp *dbProperties) DBHomePath() string {
	return dbHomePath
}

func init() {
	dbHomePath = readDBHome()
	//ex, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}
	//exPath := filepath.Dir(ex)
	//fmt.Printf("execution path: %s\n", exPath)
}

func readDBHome() string {
	home := os.Getenv(dbHomeVarName)
	if home == "" {
		panic("Env variable: " + dbHomeVarName + " is empty.")
	} else {
		log.Printf("DB home path is set to: %s\n", home)
	}

	return home
}
