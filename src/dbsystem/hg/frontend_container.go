package hg

import (
	"HomegrownDB/dbsystem/access/transaction"
	"HomegrownDB/dbsystem/auth"
)

type FrontendContainer struct {
	AuthManger         auth.Manager
	ExecutionContainer ExecutionContainer
	TxManager          transaction.Manager
}
