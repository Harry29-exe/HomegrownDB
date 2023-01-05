package testinfr

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/config"
	table2 "HomegrownDB/dbsystem/relation/table"
	"testing"
)

func NewDBStore(t *testing.T, properties config.Properties, tables ...table2.Definition) dbsystem.DBSystem {
	tableStore := table2.NewEmptyTableStore()
	for _, tableDef := range tables {
		err := tableStore.AddNewTable(tableDef)
		assert.ErrIsNil(err, t)
	}

	//ioStore := pageio.NewStore()
	//buff := buffer.NewSharedBuffer(properties.SharedBufferSize, ioStore)
	//fsmStore := fsm.NewStore()
	//_, _ = fsmStore, buff
	//for {
	//
	//}
	//todo implement me
	panic("Not implemented")
}
