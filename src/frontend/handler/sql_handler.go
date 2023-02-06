package handler

import (
	"HomegrownDB/backend"
	"HomegrownDB/dbsystem/auth"
	"HomegrownDB/dbsystem/hg"
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

func NewSqlHandler(container hg.FrontendContainer) SqlHandler {
	return stdSqlHandler{
		Container:  container.ExecutionContainer,
		AuthManger: container.AuthManger,
		TxManger:   container.TxManager,
	}
}

type stdSqlHandler struct {
	Container  hg.ExecutionContainer
	AuthManger auth.Manager
	TxManger   tx.Manager
}

var _ SqlHandler = &stdSqlHandler{}

func (s stdSqlHandler) Handle(query string, txId tx.Id, authentication auth.Authentication) (SqlResult, error) {
	user, err := s.AuthManger.Authenticate(authentication)
	if err != nil {
		return nil, err
	}
	//todo use user in backend.Execute
	_ = user

	var transaction tx.Tx
	if txId == tx.InvalidId {
		transaction = s.TxManger.New(tx.CommittedRead)
	} else {
		transaction = s.TxManger.Get(txId)
	}

	resultRows, err := backend.Execute(query, transaction, s.Container)
	if err != nil {
		return nil, err
	}
	return ToJsonResult(resultRows), nil
}

// -------------------------
//      Result
// -------------------------

var _ SqlResult = JsonResult{}

type JsonResult struct {
	data []byte
}

func (r JsonResult) Result() []byte {
	return r.data
}

func (r JsonResult) Type() SqlResultType {
	return SqlResultJSON
}
