package hg

import (
	"HomegrownDB/dbsystem/access/systable/sysinit"
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
	_ = configModule

	err = sysinit.CreateSysTables(storageModule.FS())
	if err != nil {
		return err
	}
	return nil
}
