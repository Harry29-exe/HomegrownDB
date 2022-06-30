package io

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/bstructs"
	"HomegrownDB/dbsystem/schema"
	"HomegrownDB/dbsystem/schema/table"
	"fmt"
	"os"
)

var homePath string
var tablesPaths []string     // array of table.Id to path to table data directory
var tablesDataPaths []string // array of table.Id to path to table data directory

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
		tableId := def.TableId()

		tablesPaths[tableId] = homePath + def.Name()
		tablesDataPaths[tableId] = homePath + def.Name() + "/data.bdata"
	}
}

func FetchPage(tag bstructs.PageTag, buffer []byte) {
	file, err := os.Open(tablesDataPaths[tag.TableId])
	if err != nil {
		panic(fmt.Sprintf("data file for table %s does not exist", schema.Tables.Table(tag.TableId).Name()))
	}

	_, err = file.Seek(int64(bstructs.PageSize)*int64(tag.PageId), 0)
	if err != nil {
		panic(err.Error())
	}

	_, err = file.ReadAt(buffer[:bstructs.PageSize], int64(bstructs.PageSize)*int64(tag.PageId))
	if err != nil {
		panic(err.Error())
	}
}
