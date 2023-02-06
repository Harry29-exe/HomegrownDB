package access

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/config"
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
) (relation.Manager, error) {
	manager := new(relation.Manager)
	manager.
}
