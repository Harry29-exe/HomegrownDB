package di

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/auth"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
)

func RootPathFromEnv() (string, error) {
	return config.ReadRootPathEnv()
}

func FS(rootPath string) (dbfs.FS, error) {
	return dbfs.LoadFS(rootPath)
}

func Configuration(fs dbfs.FS) (*config.Configuration, error) {
	data, err := fs.ReadConfigFile()
	if err != nil {
		return nil, err
	}
	return config.DeserializeConfig(data)
}

func Properties(fs dbfs.FS) (config.DBProperties, error) {
	return config.ReadInitProperties(fs)
}

func PageIOStore(args SimpleArgs) (pageio.Store, error) {
	return pageio.NewStore(args.FS), nil
}

func TableStore(args SimpleArgs) (table.Store, error) {
	return table.NewEmptyTableStore(), nil
}

func FsmStore(args SimpleArgs) (fsm.Store, error) {
	return fsm.NewStore(), nil
}

func SharedBuffer(args SimpleArgs, store pageio.Store) (buffer.SharedBuffer, error) {
	return buffer.NewSharedBuffer(args.C.SharedBufferSize, store), nil
}

func TxManager(args SimpleArgs) (tx.Manager, error) {
	return tx.NewManager(args.P.NextTxID), nil
}

func AuthManager(args SimpleArgs) (auth.Manager, error) {
	return auth.NewAllowAllManager(), nil
}
