package hg

import "HomegrownDB/dbsystem/hg/initializer"

type CreateArgs = initializer.CreatorProps

const (
	CreatorModeDBInitializer initializer.CreatorMode = initializer.DBInstaller
	CreatorModeTest          initializer.CreatorMode = initializer.Test
)

func Create(args CreateArgs) error {
	return initializer.InitializeDB(args)
}

func LoadFromPath(rootPath string) (DB, error) {

}

func Load() (DB, error) {
	
}