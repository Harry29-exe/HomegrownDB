package initializer

import (
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/storage/dbfs"
	"encoding/json"
)

type DBProperties struct {
	Relations []relPTR
	NextRID   relation.ID
	NextOID   dbobj.OID
}

func DefaultDBProperties() DBProperties {
	return DBProperties{
		Relations: make([]relPTR, 0),
		NextRID:   0,
		NextOID:   0,
	}
}

type relPTR struct {
	RelKind    relation.Kind
	RelationID relation.ID
}

func readInitProperties(fs dbfs.PropertiesFS) (DBProperties, error) {
	props := DBProperties{}

	fileData, err := fs.ReadPropertiesFile()

	err = json.Unmarshal(fileData, &props)
	return props, err

}

//func saveInitProperties(properties loadProps) error {
//	file, err := os.Create(config.Config.DBHomePath + "/init_props.hdb")
//	if err != nil {
//		return err
//	}
//	serialized, err := json.Marshal(properties)
//	if err != nil {
//		return err
//	}
//	_, err = file.Write(serialized)
//	if err != nil {
//		return err
//	}
//	return nil
//}
