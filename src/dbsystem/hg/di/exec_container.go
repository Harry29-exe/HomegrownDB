package di

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/fsm"
)

type ExecutionContainer struct {
	SharedBuffer buffer.SharedBuffer
	FsmStore     fsm.Store
	TableStore   table.Store
}
