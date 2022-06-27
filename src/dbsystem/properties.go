package dbsystem

import (
	"fmt"
	"os"
	"path/filepath"
)

const dbHomeVarName = "HOMEGROWN_DB_HOME"
const TableDirname = "tables"
const TableInfoFilename = "info"

var dbHomePath = ""

var Props = &dbProperties{}

type dbProperties struct{}

func init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
}

func (dbp *dbProperties) DBHomePath() string {
	return dbHomePath
}

func readDBHome() string {
	home := os.Getenv(dbHomeVarName)
	if home == "" {
		panic("Env variable: " + dbHomeVarName + " is empty.")
	}

	return home
}
