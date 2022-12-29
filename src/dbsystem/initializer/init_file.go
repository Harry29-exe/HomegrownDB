package initializer

import (
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/schema/relation"
	"encoding/json"
	"os"
)

type initProperties struct {
	Relations []relPTR
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
