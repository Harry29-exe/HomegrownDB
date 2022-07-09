package dbsystem

import (
	"log"
	"os"
)

const dbHomeVarName = "HOMEGROWN_DB_HOME"
const TableDirname = "tables"
const TableInfoFilename = "info"

var dbHomePath = ""

func init() {
	dbHomePath = readDBHome()
}

// DBHomePath returns path to root directory of database without '/' postfix
func DBHomePath() string {
	return dbHomePath
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
