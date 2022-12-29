package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

// todo init
// DBHomePath returns path to root directory of database without '/' postfix
var DBHomePath string
var Props *Properties

func init() {
	DBHomePath, _ = readDBHome()
	var err error
	Props, err = ReadConfig()
	if err != nil {
		panic(err.Error())
	}
}

func ReadConfig() (*Properties, error) {
	dbHomePath, err := readDBHome()
	if err != nil {
		return nil, err
	}

	return readConfigFile(dbHomePath)
}

func readDBHome() (string, error) {
	home := os.Getenv(dbHomeVarName)
	if home == "" {
		return home, errors.New("env variable: " + dbHomeVarName + " is empty")
	} else {
		log.Printf("DB home path is set to: %s\n", home)
	}

	return home, nil
}

func readConfigFile(dbHomePath string) (*Properties, error) {
	conf := &Properties{}

	file, err := os.Open(dbHomePath + "/config.hdb")
	if err != nil {
		return conf, err
	}
	fileStats, err := file.Stat()
	if err != nil {
		return conf, err
	}

	fileData := make([]byte, fileStats.Size())
	_, err = file.Read(fileData)
	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(fileData, conf)
	return conf, err
}
