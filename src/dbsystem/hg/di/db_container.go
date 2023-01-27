package di

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/auth"
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
	AuthManager  auth.Manager
}

func (c *Container) CreateFrontendContainer() FrontendContainer {
	return FrontendContainer{
		AuthManger:         c.AuthManager,
		ExecutionContainer: c.CreateExecutionContainer(),
		TxManager:          c.TxManager,
	}
}

func (c *Container) CreateExecutionContainer() ExecutionContainer {
	return ExecutionContainer{
		SharedBuffer: c.SharedBuffer,
		FsmStore:     c.FsmStore,
		TableStore:   c.TableStore,
	}
}
