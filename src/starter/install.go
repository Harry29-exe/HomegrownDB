package starter

import (
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/dbsystem/hglib"
	"os"
)

func Install(rootPath string) error {
	err := hg.CreateDB(hglib.ModuleInstallerArgs{
		Mode:     hglib.InstallerModeDB,
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
