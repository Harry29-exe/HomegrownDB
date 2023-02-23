package config

import "HomegrownDB/dbsystem/storage/dbfs"

func ConfigProvider(fs dbfs.FS) (*Configuration, error) {
	data, err := fs.ReadConfigFile()
	if err != nil {
		return nil, err
	}
	return DeserializeConfig(data)
}
