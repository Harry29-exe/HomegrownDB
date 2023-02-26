package access

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/access/sysinit"
	"HomegrownDB/dbsystem/hglib"
)

func InstallModule(mode hglib.Mode, deps ModuleDeps) (Module, error) {
	buffSize := deps.ConfigModule.Config().SharedBufferSize
	buff := buffer.NewSharedBuffer(buffSize, deps.StorageModule.PageIOStore())

	err := sysinit.Phase1.Execute(deps.StorageModule.FS())
	if err != nil {
		return nil, err
	}
	relManager, err := relation.NewManager(buff, deps.StorageModule.FS())
	if err != nil {
		return nil, err
	}
	err = sysinit.Phase2.Execute(relManager)
	if err != nil {
		return nil, err
	}
	err = sysinit.Phase3.Execute(relManager)
	if err != nil {
		return nil, err
	}

	err = buff.FlushAll()
	if err != nil {
		return nil, err
	}

	txManager, err := TxManagerProvider(deps.ConfigModule, buff)
	if err != nil {
		return nil, err
	}

	return &stdModule{
		sharedBuffer:    buff,
		relationManager: relManager,
		txManager:       txManager,
	}, nil
}
