package hg

import (
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/hg/internal/creator"
	"HomegrownDB/dbsystem/storage"
)

type CreateArgs = creator.Props

const (
	CreatorModeDBInitializer creator.Mode = creator.DBInstaller
	CreatorModeTest          creator.Mode = creator.Test
)

func Create(args CreateArgs) error {
	return creator.CreateDB(args)
}

// Load create DB object with provided FutureContainer, if fc is nil then Load
// will create default
func Load(builders *MBuilders) (DB, error) {
	if builders == nil {
		builders = DefaultMBuilders()
	}
	var err error
	dbModule := new(DBSystem)

	dbModule.storageModule, err = storage.NewModule(builders.StorageMBuilder)
	if err != nil {
		return nil, err
	}

	dbModule.configModule, err = config.NewModule(
		builders.ConfigMBuilder,
		config.ModuleDeps{StorageModule: dbModule.storageModule})
	if err != nil {
		return nil, err
	}

	dbModule.accessModule, err = access.NewModule(builders.AccessMBuilder, access.ModuleDeps{
		StorageModule: dbModule.storageModule,
		ConfigModule:  dbModule.configModule,
	})
	if err != nil {
		return nil, err
	}

	return dbModule, nil
}

// MBuilders wrapper to keep all module builders
type MBuilders struct {
	StorageMBuilder storage.ModuleBuilder
	ConfigMBuilder  config.ModuleBuilder
	AccessMBuilder  access.ModuleBuilder
}

func DefaultMBuilders() *MBuilders {
	return &MBuilders{
		StorageMBuilder: storage.DefaultModuleBuilder(),
		ConfigMBuilder:  config.DefaultBuilder(),
		AccessMBuilder:  access.DefaultModuleBuilder(),
	}
}
