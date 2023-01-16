package hg

import "HomegrownDB/dbsystem/hg/di"

type CreateArgs = CreatorProps

const (
	CreatorModeDBInitializer CreatorMode = DBInstaller
	CreatorModeTest          CreatorMode = Test
)

func Create(args CreateArgs) error {
	return CreateDB(args)
}

func LoadFromPath(rootPath string) (DB, error) {
	//todo implement me
	panic("Not implemented")
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
	}
}
