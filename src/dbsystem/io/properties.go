package io

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/stores"
)

var homePath string
var tablesPaths []string     // array of table.Id to path to table data directory
var tablesDataPaths []string // array of table.Id to path to table data directory

func init() {
	homePath = dbsystem.DBHomePath() + "/tables"
	initIOProperties()
	stores.Tables.RegisterChangeListener(initIOProperties)
}

func initIOProperties() {
	tables := stores.Tables.AllTables()

	maxTableId := table.Id(0)
	for _, def := range tables {
		if def.TableId() > maxTableId {
			maxTableId = def.TableId()
		}
	}

	tablesPaths = make([]string, maxTableId)
	for _, def := range tables {
		tableId := def.TableId()

		tablesPaths[tableId] = homePath + def.Name()
		tablesDataPaths[tableId] = homePath + def.Name() + "/data.bdata"
	}
}
