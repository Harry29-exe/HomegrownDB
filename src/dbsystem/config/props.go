package config

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
