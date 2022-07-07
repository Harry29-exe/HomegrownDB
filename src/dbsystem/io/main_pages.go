package io

import (
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/stores"
	"fmt"
	"os"
)

var Pages = &PageIO{}

type PageIO struct {
	homePath        string
	tablesPaths     []string // array of table.Id to path to table data directory
	tablesDataPaths []string // array of table.Id to path to table data directory
}

func (mp *PageIO) Read(tag bdata.PageTag, buffer []byte) {
	file, err := os.Open(mp.tablesDataPaths[tag.TableId])
	if err != nil {
		panic(fmt.Sprintf("data file for table %s does not exist", stores.Tables.Table(tag.TableId).Name()))
	}

	_, err = file.ReadAt(buffer[:bdata.PageSize], int64(bdata.PageSize)*int64(tag.PageId))
	if err != nil {
		panic(err.Error())
	}
}

func (mp *PageIO) Flush(tag bdata.PageTag, buffer []byte) {
	//todo implement me
	panic("Not implemented")
}

func (mp *PageIO) SaveNew(tag bdata.PageTag, buffer []byte) {
	//todo implement me
	panic("Not implemented")
}
