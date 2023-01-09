package hg

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

func Load() (DB, error) {
	//todo implement me
	panic("Not implemented")
}
