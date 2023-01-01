package config

import (
	"HomegrownDB/dbsystem/config/envvar"
	"errors"
	"log"
	"os"
)

func ReadRootPathEnv() (string, error) {
	home := os.Getenv(dbHomeVarName)
	if home == "" {
		return home, errors.New("env variable: " + dbHomeVarName + " is empty")
	} else {
		log.Printf("DB home path is set to: %s\n", home)
	}

	return home, nil
}

func SetRootPathEnv(rootPath string) error {
	return envvar.SetOsEnv(dbHomeVarName, rootPath)
}
