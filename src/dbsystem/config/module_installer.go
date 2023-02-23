package config

import (
	"HomegrownDB/dbsystem/hglib"
	"fmt"
	"log"
)

func InstallModule(args hglib.ModuleInstallerArgs, deps ModuleDeps) (Module, error) {
	if args.RootPath == "" {
		err := fmt.Errorf("illegal root path %s", args.RootPath)
		log.Printf(err.Error())
		return nil, err
	}

	config := &Configuration{
		DBHomePath:       args.RootPath,
		SharedBufferSize: getBufferSize(args),
	}

	err := deps.StorageModule.FS().InitDBSystemConfig(SerializeConfig(*config))
	if err != nil {
		return nil, err
	}

	return NewModule(DefaultBuilder(), deps)
}

const (
	DefaultBufferSize = 10_000
)

func getBufferSize(args hglib.ModuleInstallerArgs) uint {
	if args.BufferSize == 0 {
		return DefaultBufferSize
	}
	return args.BufferSize
}
