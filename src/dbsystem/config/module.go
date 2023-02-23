package config

import (
	"HomegrownDB/dbsystem/storage"
	"HomegrownDB/dbsystem/storage/dbfs"
)

type Module interface {
	Config() *Configuration
}

type ModuleBuilder struct {
	ConfigProvider func(fs dbfs.FS) (*Configuration, error)
}

func DefaultBuilder() ModuleBuilder {
	return ModuleBuilder{
		ConfigProvider: ConfigProvider,
	}
}

type ModuleDeps struct {
	StorageModule storage.Module
}

func NewModule(builder ModuleBuilder, deps ModuleDeps) (Module, error) {
	module := new(stdModule)
	var err error
	module.config, err = builder.ConfigProvider(deps.StorageModule.FS())
	if err != nil {
		return nil, err
	}

	return module, nil
}

// -------------------------
//      internal
// -------------------------

type stdModule struct {
	config *Configuration
}

func (s *stdModule) Config() *Configuration {
	return s.config
}
