package storage

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/storage/dbfs"
)

func InstallModule(args hglib.ModuleInstallerArgs) (Module, error) {
	if args.Mode == hglib.InstallerModeDB {
		err := SetRootPathEnv(args.RootPath)
		if err != nil {
			return nil, err
		}
	}

	fs, err := dbfs.CreateFS(args.RootPath)
	if err != nil {
		return nil, err
	}
	err = fs.InitDBSystemDirs()
	if err != nil {
		return nil, err
	}

	modBuilder := DefaultModuleBuilder()
	modBuilder.RootPathProvider = func() (string, error) { return args.RootPath, nil }

	return NewModule(modBuilder)
}
