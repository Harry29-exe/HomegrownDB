package di

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
)

func RootPathFromEnv() (string, error) {
	return config.ReadRootPathEnv()
}

func FS(rootPath string) (dbfs.FS, error) {
	return dbfs.LoadFS(rootPath)
}

func Configuration(fs dbfs.FS) (*config.Configuration, error) {
	return config.ReadConfig(fs)
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
