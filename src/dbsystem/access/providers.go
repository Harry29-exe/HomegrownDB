package access

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/access/transaction"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/storage"
	"HomegrownDB/dbsystem/tx"
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
	return relation.NewManager(buff, storageModule.FS())
}

func TxManagerProvider(
	configModule config.Module,
	buffer buffer.SharedBuffer,
) (transaction.Manager, error) {
	return transaction.NewManager(tx.Id(1000)), nil
}
