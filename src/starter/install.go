package starter

import (
	"HomegrownDB/dbsystem/hg"
	"os"
)

func Install(rootPath string) error {
	err := hg.Create(hg.CreateArgs{
		Mode:     hg.CreatorModeDBInitializer,
		RootPath: rootPath,
	})

	return err
}

func InstallDefault() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dbPath := homeDir + "/" + ".HomeGrownDB"
	return Install(dbPath)
}
