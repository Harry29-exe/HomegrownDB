package starter

import (
	"HomegrownDB/dbsystem/storage"
	"os"
)

func UninstallDefault() error {
	path, err := storage.ReadRootPathEnv()
	if err != nil {
		return err
	}
	err = os.RemoveAll(path)
	if err != nil {
		return err
	}
	return storage.ClearRootPathEnv()
}
