package di

import "HomegrownDB/dbsystem/auth"

type FrontendContainer struct {
	AuthManger         auth.Manager
	ExecutionContainer ExecutionContainer
}
