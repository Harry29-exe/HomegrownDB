package testinfr

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
	"testing"
)

func NewDBStore(t *testing.T, properties config.Properties, tables ...table.Definition) dbsystem.DBSystem {
	tableStore := table.NewEmptyTableStore()
	for _, tableDef := range tables {
		err := tableStore.AddNewTable(tableDef)
		assert.ErrIsNil(err, t)
	}

	ioStore := pageio.NewStore()
	buff := buffer.NewSharedBuffer(properties.SharedBufferSize, ioStore)
	fsmStore := fsm.NewStore(buff)
	_ = fsmStore
	//for {
	//
	//}
	//todo implement me
	panic("Not implemented")
}
