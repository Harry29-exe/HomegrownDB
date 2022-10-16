package dbfs

import "HomegrownDB/dbsystem/schema/table"

func PathToTableDir(table table.Definition) string {
	return DBHomePath + TablesDirPath + "/" + table.Name()
}
