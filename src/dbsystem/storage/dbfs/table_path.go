package dbfs

import "HomegrownDB/dbsystem/schema/table"

func PathToTableDir(table table.RDefinition) string {
	return DBHomePath + TablesDirPath + "/" + table.Name()
}
