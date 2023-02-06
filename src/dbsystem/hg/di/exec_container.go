package di

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/storage/fsm"
)

type ExecutionContainer struct {
	SharedBuffer    buffer.SharedBuffer
	FsmStore        fsm.Store
	RelationManager relation.Manager
}
