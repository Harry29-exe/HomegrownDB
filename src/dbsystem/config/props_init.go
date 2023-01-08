package config

import (
	"HomegrownDB/dbsystem/config/envvar"
	"HomegrownDB/dbsystem/storage/dbfs"
	"encoding/json"
	"errors"
	"log"
	"os"
)

func ReadConfig(fs dbfs.PropertiesFS) (*Configuration, error) {
	fileData, err := fs.ReadConfigFile()
	if err != nil {
		return nil, err
	}
	conf := &Configuration{}
	err = json.Unmarshal(fileData, conf)

	return conf, err
}

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
