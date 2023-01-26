package di

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
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
	TxManager    tx.Manager
}

func (c *Container) CreateFrontendContainer() FrontendContainer {
	return FrontendContainer{
		AuthManger:         nil,
		ExecutionContainer: c.CreateExecutionContainer(),
	}
}

func (c *Container) CreateExecutionContainer() ExecutionContainer {
	return ExecutionContainer{
		SharedBuffer: c.SharedBuffer,
		FsmStore:     c.FsmStore,
		TableStore:   c.TableStore,
	}
}
