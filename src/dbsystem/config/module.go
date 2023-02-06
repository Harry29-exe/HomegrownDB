package config

import (
	"HomegrownDB/dbsystem/storage"
	"HomegrownDB/dbsystem/storage/dbfs"
)

type Module interface {
	Config() *Configuration
	Props() DBProperties
}

type ModuleBuilder struct {
	ConfigProvider    func(fs dbfs.FS) (*Configuration, error)
	PropertiesProvide func(fs dbfs.FS) (DBProperties, error)
}

func DefaultBuilder() ModuleBuilder {
	return ModuleBuilder{
		ConfigProvider:    ConfigProvider,
		PropertiesProvide: PropertiesProvider,
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
	module.props, err = builder.PropertiesProvide(deps.StorageModule.FS())
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
	props  DBProperties
}

func (s *stdModule) Config() *Configuration {
	return s.config
}

func (s *stdModule) Props() DBProperties {
	return s.props
}
