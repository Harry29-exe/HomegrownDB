package initializer

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
)

func createCtx(props *config.Properties, initProps initProperties) *ctx {
	return &ctx{
		Props:       props,
		InitProps:   initProps,
		PageIOStore: pageio.NewStore(),
		TableStore:  table.NewEmptyTableStore(),
	}
}

type ctx struct {
	Props     *config.Properties
	InitProps initProperties

	FS           dbfs.FS
	PageIOStore  *pageio.Store
	SharedBuffer buffer.SharedBuffer
	TableStore   table.Store
	FsmStore     fsm.Store
}
