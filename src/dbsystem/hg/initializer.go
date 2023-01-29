package hg

import (
	"HomegrownDB/dbsystem/hg/di"
	"HomegrownDB/dbsystem/hg/internal/creator"
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
func Load(fc *di.FutureContainer) (DB, error) {
	if fc == nil {
		fcTemp := DefaultFutureContainer()
		fc = &fcTemp
	}

	container, err := fc.Build()
	if err != nil {
		return nil, err
	}
	return NewDB(container), nil
}

func DefaultFutureContainer() di.FutureContainer {
	return di.FutureContainer{
		RootProvider:         di.RootPathFromEnv,
		FsProvider:           di.FS,
		ConfigProvider:       di.Configuration,
		PropertiesProvider:   di.Properties,
		PageIOStoreProvider:  di.PageIOStore,
		TableStoreProvider:   di.TableStore,
		FsmStoreProvider:     di.FsmStore,
		SharedBufferProvider: di.SharedBuffer,
		TxManagerProvider:    di.TxManager,
		AuthManagerProvider:  di.AuthManager,
	}
}
