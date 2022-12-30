package main

import (
	"encoding/json"
	"os"
)

func main() {
	env, ok := os.LookupEnv("HOMEGROWN_DB_HOME")
	if !ok {
		os.Exit(1)
	}
	file, err := os.Create(env + "/config.hdb")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	props := Properties{
		DBHomePath:       env,
		SharedBufferSize: 1000,
	}
	propsJson, err := json.Marshal(props)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	_, err = file.Write(propsJson)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

type Properties struct {
	DBHomePath       string
	SharedBufferSize uint
}

