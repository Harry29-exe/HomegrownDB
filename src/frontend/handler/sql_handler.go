package handler

import (
	"HomegrownDB/dbsystem/auth"
	"HomegrownDB/dbsystem/tx"
)

type SqlHandler interface {
	Handle(query string, tx tx.Id, auth auth.Authentication) (SqlResult, error)
}

type SqlResult interface {
	Result() []byte
	Type() SqlResultType
}

type SqlResultType int8

const (
	SqlResultBinary SqlResultType = iota
	SqlResultJSON
)
