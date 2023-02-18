package config

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/tx"
	"encoding/json"
)

func DeserializeProperties(propsData []byte) (DBProperties, error) {
	props := DBProperties{}
	err := json.Unmarshal(propsData, &props)
	return props, err
}

func SerializeProperties(properties DBProperties) []byte {
	data, err := json.Marshal(properties)
	if err != nil {
		panic(err.Error())
	}
	return data
}

type DBProperties struct {
	NextTxID tx.Id
	NextOID  hglib.OID
}

func DefaultDBProperties() DBProperties {
	return DBProperties{
		NextOID: 10_000, // OIDs less than 10_000 are reserved for system tables
	}
}

func ReadInitProperties(fs dbfs.PropertiesFS) (DBProperties, error) {
	props := DBProperties{}

	fileData, err := fs.ReadPropertiesFile()

	err = json.Unmarshal(fileData, &props)
	return props, err

}
