package config

import "encoding/json"

func DeserializeConfig(configData []byte) (*Configuration, error) {
	conf := &Configuration{}
	err := json.Unmarshal(configData, conf)

	return conf, err
}

func SerializeConfig(config Configuration) []byte {
	data, err := json.Marshal(config)
	if err != nil {
		panic(err.Error())
	}
	return data
}

type Configuration struct {
	DBHomePath       string
	SharedBufferSize uint
}

func DefaultConfiguration(rootPath string) Configuration {
	return Configuration{
		DBHomePath:       rootPath,
		SharedBufferSize: 10_000,
	}
}
