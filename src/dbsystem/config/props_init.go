package config

import (
	"HomegrownDB/dbsystem/storage/dbfs"
	"encoding/json"
	"errors"
	"log"
	"os"
)

func ReadConfig(fs dbfs.PropertiesFS) (*Properties, error) {
	fileData, err := fs.ReadConfigFile()
	if err != nil {
		return nil, err
	}
	conf := &Properties{}
	err = json.Unmarshal(fileData, conf)

	return conf, err
}

func ReadRootPath() (string, error) {
	home := os.Getenv(dbHomeVarName)
	if home == "" {
		return home, errors.New("env variable: " + dbHomeVarName + " is empty")
	} else {
		log.Printf("DB home path is set to: %s\n", home)
	}

	return home, nil
}
