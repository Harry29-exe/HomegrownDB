package storage

import (
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/pageio"
)

type Module interface {
	PageIOStore() pageio.Store
	FS() dbfs.FS
}

type ModuleBuilder struct {
	RootPathProvider    func() (string, error)
	FsProvider          func(rootPath string) (dbfs.FS, error)
	PageIOStoreProvider func(fs dbfs.FS) (pageio.Store, error)
}

func DefaultModuleBuilder() ModuleBuilder {
	return ModuleBuilder{
		RootPathProvider: ReadRootPathEnv,
		FsProvider:       dbfs.LoadFS,
		PageIOStoreProvider: func(fs dbfs.FS) (pageio.Store, error) {
			return pageio.NewStore(fs), nil
		},
	}
}

func NewModule(dependencies ModuleBuilder) (Module, error) {
	module := new(stdModule)
	rootPath, err := dependencies.RootPathProvider()
	if err != nil {
		return nil, err
	}
	module.fs, err = dependencies.FsProvider(rootPath)
	if err != nil {
		return nil, err
	}
	module.pageIOStore, err = dependencies.PageIOStoreProvider(module.fs)
	if err != nil {
		return nil, err
	}

	return module, nil
}

// -------------------------
//      internal
// -------------------------

type stdModule struct {
	pageIOStore pageio.Store
	fs          dbfs.FS
}

func (s *stdModule) PageIOStore() pageio.Store {
	return s.pageIOStore
}

func (s *stdModule) FS() dbfs.FS {
	return s.fs
}
