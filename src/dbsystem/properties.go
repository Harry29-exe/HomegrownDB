package dbsystem

import "os"

const dbHomeVarName = "HOMEGROWN_DB_HOME"
const TableDirname = "tables"
const TableInfoFilename = "info"

var dbHomePath = ""

//readDBHome()

func GetDBHomePath() string {
	return dbHomePath
}

func readDBHome() string {
	home := os.Getenv(dbHomeVarName)
	if home == "" {
		panic("Env variable: " + dbHomeVarName + " is empty.")
	}

	return home
}
