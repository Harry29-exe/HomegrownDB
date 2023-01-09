package config

import (
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/storage/dbfs"
	"encoding/json"
)

type DBProperties struct {
	Relations []RelPTR
	NextRID   relation.ID
	NextOID   dbobj.OID
}

func DefaultDBProperties() DBProperties {
	return DBProperties{
		Relations: make([]RelPTR, 0),
		NextRID:   0,
		NextOID:   0,
	}
}

type RelPTR struct {
	RelKind    relation.Kind
	RelationID relation.ID
}

func ReadInitProperties(fs dbfs.PropertiesFS) (DBProperties, error) {
	props := DBProperties{}

	fileData, err := fs.ReadPropertiesFile()

	err = json.Unmarshal(fileData, &props)
	return props, err

}
