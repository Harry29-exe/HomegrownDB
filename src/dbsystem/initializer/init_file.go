package initializer

import (
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/schema/dbobj"
	"HomegrownDB/dbsystem/schema/relation"
	"encoding/json"
	"os"
)

type initProperties struct {
	Relations []relPTR
	NextRID   relation.ID
	NextOID   dbobj.OID
}

type relPTR struct {
	RelKind relation.Kind
	// RelativePath is path to serialized definition relative to DBHomePath
	RelativePath string
}

func readInitProperties() (initProperties, error) {
	props := initProperties{}

	file, err := os.Open(config.Props.DBHomePath + "/init_props.hdb")
	if err != nil {
		return props, err
	}
	fileStats, err := file.Stat()
	if err != nil {
		return props, err
	}

	fileData := make([]byte, fileStats.Size())
	_, err = file.Read(fileData)
	if err != nil {
		return props, err
	}

	err = json.Unmarshal(fileData, &props)
	return props, err

}

func saveInitProperties(properties initProperties) error {
	file, err := os.Create(config.Props.DBHomePath + "/init_props.hdb")
	if err != nil {
		return err
	}
	serialized, err := json.Marshal(properties)
	if err != nil {
		return err
	}
	_, err = file.Write(serialized)
	if err != nil {
		return err
	}
	return nil
}
