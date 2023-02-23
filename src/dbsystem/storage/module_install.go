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

	return NewModule(DefaultModuleBuilder())
}
