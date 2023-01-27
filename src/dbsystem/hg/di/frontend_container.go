package di

import (
	"HomegrownDB/dbsystem/auth"
	"HomegrownDB/dbsystem/tx"
)

type FrontendContainer struct {
	AuthManger         auth.Manager
	ExecutionContainer ExecutionContainer
	TxManager          tx.Manager
}
