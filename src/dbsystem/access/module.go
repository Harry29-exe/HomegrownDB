package access

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/storage"
)

type Module interface {
	SharedBuffer() buffer.SharedBuffer
	RelationManager() relation.Manager
}

type ModuleBuilder struct {
	StorageModule storage.Module
	ConfigModule  config.Module

	SharedBufferProvider   func(storageModule storage.Module, configModule config.Module) (buffer.SharedBuffer, error)
	RelationMangerProvider func(module storage.Module, buff buffer.SharedBuffer) (relation.Manager, error)
}

func DefaultModuleBuilder(configModule config.Module, storageModule storage.Module) ModuleBuilder {
	return ModuleBuilder{
		StorageModule:          storageModule,
		ConfigModule:           configModule,
		SharedBufferProvider:   SharedBufferProvider,
		RelationMangerProvider: RelationManagerProvider,
	}
}

func NewModule(builder ModuleBuilder) (Module, error) {
	var err error
	module := new(stdModule)

	module.sharedBuffer, err = builder.SharedBufferProvider(builder.StorageModule, builder.ConfigModule)
	if err != nil {
		return nil, err
	}
	module.relationManager, err = builder.RelationMangerProvider(builder.StorageModule, module.sharedBuffer)
	if err != nil {
		return nil, err
	}

	return module, nil
}

// -------------------------
//      internal
// -------------------------

type stdModule struct {
	sharedBuffer    buffer.SharedBuffer
	relationManager relation.Manager
}

func (s *stdModule) SharedBuffer() buffer.SharedBuffer {
	return s.sharedBuffer
}

func (s *stdModule) RelationManager() relation.Manager {
	return s.relationManager
}
