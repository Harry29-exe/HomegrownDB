package initializer

import (
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/schema/table"
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

	PageIOStore *pageio.Store
	TableStore  table.Store
}
