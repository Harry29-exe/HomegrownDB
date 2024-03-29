package access

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/access/sequence"
	"HomegrownDB/dbsystem/access/transaction"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage"
)

type Module interface {
	hglib.Module
	SharedBuffer() buffer.SharedBuffer
	RelationManager() relation.Manager
	TxManager() transaction.Manager
}

type ModuleBuilder struct {
	SharedBufferProvider   func(storageModule storage.Module, configModule config.Module) (buffer.SharedBuffer, error)
	OIDSequenceProvider    func(module config.Module) (sequence.Sequence[reldef.OID], error)
	RelationMangerProvider func(module storage.Module, buff buffer.SharedBuffer) (relation.Manager, error)
	TxManagerProvider      func(configModule config.Module, buffer buffer.SharedBuffer) (transaction.Manager, error)
}

func DefaultModuleBuilder() ModuleBuilder {
	return ModuleBuilder{
		SharedBufferProvider:   SharedBufferProvider,
		RelationMangerProvider: RelationManagerProvider,
		TxManagerProvider:      TxManagerProvider,
	}
}

type ModuleDeps struct {
	StorageModule storage.Module
	ConfigModule  config.Module
}

func NewModule(builder ModuleBuilder, deps ModuleDeps) (Module, error) {
	var err error
	module := new(stdModule)

	module.sharedBuffer, err = builder.SharedBufferProvider(deps.StorageModule, deps.ConfigModule)
	if err != nil {
		return nil, err
	}

	module.relationManager, err = builder.RelationMangerProvider(deps.StorageModule, module.sharedBuffer)
	if err != nil {
		return nil, err
	}

	module.txManager, err = builder.TxManagerProvider(deps.ConfigModule, module.sharedBuffer)

	return module, nil
}

// -------------------------
//      internal
// -------------------------

type stdModule struct {
	sharedBuffer    buffer.SharedBuffer
	relationManager relation.Manager
	txManager       transaction.Manager
}

func (s *stdModule) SharedBuffer() buffer.SharedBuffer {
	return s.sharedBuffer
}

func (s *stdModule) RelationManager() relation.Manager {
	return s.relationManager
}

func (s *stdModule) TxManager() transaction.Manager {
	return s.txManager
}

func (s *stdModule) Shutdown() error {
	return s.sharedBuffer.FlushAll()
}
