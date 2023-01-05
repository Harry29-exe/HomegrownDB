package initializer

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/config"
	table2 "HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
)

func createCtx(props *config.Properties, initProps initProperties) *LoaderCtx {
	return &LoaderCtx{
		Config:    props,
		InitProps: initProps,
	}
}

type LoaderCtx struct {
	RootPath string
	FS       dbfs.FS

	Config    *config.Properties
	InitProps initProperties

	PageIOStore  *pageio.StdStore
	SharedBuffer buffer.SharedBuffer
	TableStore   table2.Store
	FsmStore     fsm.Store
}
