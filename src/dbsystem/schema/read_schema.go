package schema

import (
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/schema/table"
	"os"
)

func ReadSchema(tableDirPath string) ([]table.WDefinition, error) {
	tableDir, err := os.Open(tableDirPath)

	if err != nil {
		return nil, err
	}
	tableDirs, err := tableDir.ReadDir(0)
	tableArray := make([]table.WDefinition, len(tableDirs))
	for i, dir := range tableDirs {
		if !dir.IsDir() {
			continue
		}
		tableInfoFilepath := tableDirPath + "/" + dir.Name() + config.TableDefinition
		tableInfoFile, err := os.Open(tableInfoFilepath)
		if err != nil {
			return nil, err
		}

		stat, err := tableInfoFile.Stat()
		if err != nil {
			return nil, err
		}
		fLength := stat.Size()
		buffer := make([]byte, fLength)
		_, err = tableInfoFile.Read(buffer)
		if err != nil {
			return nil, err
		}

		tableDef := table.Deserialize(buffer)
		tableArray[i] = tableDef
	}

	return tableArray, nil
}
