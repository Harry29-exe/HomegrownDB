package dbsystem

import "os"

const dbHomeVarName = "HOMEGROWN_DB_HOME"
const TableDirname = "tables"
const TableInfoFilename = "info"

var dbHomePath string = readDBHome()

func GetDBHomePath() string {
	return dbHomePath
}

func readDBHome() string {
	home := os.Getenv(dbHomeVarName)
	if home == "" {
		panic("Env variable: " + dbHomeVarName + " is empty.")
	}

	dbHomePath = home
	return home
}
