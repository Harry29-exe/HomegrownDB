package di

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
)

type Container struct {
	RootPath string
	FS       dbfs.FS
	Config   *config.Configuration
	DBProps  config.DBProperties

	PageIOStore pageio.Store
	TableStore  table.Store
	FsmStore    fsm.Store

	SharedBuffer buffer.SharedBuffer
}
