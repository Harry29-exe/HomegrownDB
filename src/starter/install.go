package starter

import (
	"HomegrownDB/dbsystem/hg"
	"os"
)

func Install(rootPath string) error {
	err := hg.CreateDB(hg.CreateArgs{
		Mode:     hg.InstallerModeDB,
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
