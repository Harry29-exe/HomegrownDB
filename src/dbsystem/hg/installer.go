package hg

import (
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/storage"
)

func CreateDB(args hglib.ModuleInstallerArgs) error {
	storageModule, err := storage.InstallModule(args)
	if err != nil {
		return err
	}

	configModule, err := config.InstallModule(args, config.ModuleDeps{StorageModule: storageModule})
	if err != nil {
		return err
	}

	accessModule, err := access.InstallModule(args.Mode, access.ModuleDeps{
		StorageModule: storageModule,
		ConfigModule:  configModule,
	})
	//err = sysinit.createSysTables(storageModule.FS())
	//if err != nil {
	//	return err
	//}
	_ = accessModule
	if err != nil {
		return err
	}

	return nil
}
