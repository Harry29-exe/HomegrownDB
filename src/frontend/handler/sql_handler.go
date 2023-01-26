package handler

import (
	"HomegrownDB/backend"
	"HomegrownDB/dbsystem/auth"
	"HomegrownDB/dbsystem/hg/di"
	"HomegrownDB/dbsystem/tx"
)

type SqlHandler interface {
	Handle(query string, txId tx.Id, auth auth.Authentication) (SqlResult, error)
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

type stdSqlHandler struct {
	Container  di.ExecutionContainer
	AuthManger auth.Manager
	TxManger   tx.Manager
}

var _ SqlHandler = &stdSqlHandler{}

func (s stdSqlHandler) Handle(query string, txId tx.Id, authentication auth.Authentication) (SqlResult, error) {
	user, err := s.AuthManger.Authenticate(authentication)
	if err != nil {
		return nil, err
	}

	var transaction tx.Tx
	if txId == tx.InvalidId {
		transaction = s.TxManger.New(tx.CommittedRead)
	} else {
		transaction = s.TxManger.Get(txId)
	}

	resultRows, err := backend.Execute(query, transaction, s.Container)

}
