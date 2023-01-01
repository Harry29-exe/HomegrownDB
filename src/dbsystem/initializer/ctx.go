package initializer

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/config"
	table2 "HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
)

func createCtx(props *config.Properties, initProps initProperties) *ctx {
	return &ctx{
		Props:       props,
		InitProps:   initProps,
		PageIOStore: pageio.NewStore(),
		TableStore:  table2.NewEmptyTableStore(),
	}
}

type ctx struct {
	Props     *config.Properties
	InitProps initProperties

	FS           dbfs.FS
	PageIOStore  *pageio.StdStore
	SharedBuffer buffer.SharedBuffer
	TableStore   table2.Store
	FsmStore     fsm.Store
}
