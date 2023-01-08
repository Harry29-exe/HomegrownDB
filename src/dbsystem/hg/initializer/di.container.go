package initializer

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
)

type Container struct {
	Tables   table.Store
	FSMs     fsm.Store
	PageIO   pageio.Store
	DBBuffer buffer.SharedBuffer
	FS       dbfs.FS
}
