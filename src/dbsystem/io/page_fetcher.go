package io

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/bstructs"
	"HomegrownDB/dbsystem/schema"
	"HomegrownDB/dbsystem/schema/table"
)

var homePath string
var tablesPaths []string // array of table.Id to path to table data directory

func init() {
	homePath = dbsystem.Props.DBHomePath() + "/tables"
	initPageFetcher()
	schema.Tables.RegisterChangeListener(initPageFetcher)
}

func initPageFetcher() {
	tables := schema.Tables.AllTables()

	maxTableId := table.Id(0)
	for _, def := range tables {
		if def.TableId() > maxTableId {
			maxTableId = def.TableId()
		}
	}

	tablesPaths = make([]string, maxTableId)
	for _, def := range tables {
		tablesPaths[def.TableId()] = homePath + def.Name()
	}
}

func FetchPage(tag bstructs.PageTag) {

	//tablesPaths[tag.TableId]
}
