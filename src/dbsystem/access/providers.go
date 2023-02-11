package access

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/access/sequence"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage"
)

func SharedBufferProvider(
	storageModule storage.Module,
	configModule config.Module,
) (buffer.SharedBuffer, error) {
	buff := buffer.NewSharedBuffer(configModule.Config().SharedBufferSize, storageModule.PageIOStore())
	return buff, nil
}

func RelationManagerProvider(
	storageModule storage.Module,
	buff buffer.SharedBuffer,
	oidSequence relation.OIDSequence,
) (relation.Manager, error) {
	return relation.NewManager(buff, storageModule.FS(), oidSequence)
}

func OIDSequenceProvider(configModule config.Module) (sequence.Sequence[reldef.OID], error) {
	return sequence.NewInMemSequence(configModule.Props().NextOID), nil
}
