package io

import (
	"HomegrownDB/dbsystem/bstructs"
	"HomegrownDB/dbsystem/schema"
	"fmt"
	"os"
)

var Pages = mainPages{}

type mainPages struct{}

func (mp mainPages) Read(tag bstructs.PageTag, buffer []byte) {
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

func (mp mainPages) Flush(tag bstructs.PageTag, buffer []byte) {
	panic("Not implemented")
}
