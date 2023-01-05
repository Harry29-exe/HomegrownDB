package initializer

import (
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/storage/dbfs"
	"encoding/json"
)

type initProperties struct {
	Relations []relPTR
	NextRID   relation.ID
	NextOID   dbobj.OID
}

type relPTR struct {
	RelKind    relation.Kind
	RelationID relation.ID
}

func readInitProperties(fs dbfs.PropertiesFS) (initProperties, error) {
	props := initProperties{}

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
