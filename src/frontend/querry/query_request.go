package querry

import (
	"HomegrownDB/backend"
	"HomegrownDB/dbsystem/stores"
	"HomegrownDB/dbsystem/tx"
	"io"
	"strings"
)

type DBRequest struct {
	txId  tx.Id
	query string
	user  any //not in use for now
}

type DBResponse struct {
	Body   io.Reader
	Status responseStatus
}

func (r *DBRequest) Handle() *DBResponse {
	var txCtx tx.Ctx
	if r.txId == 0 {
		txCtx = stores.DBTxStore.NewCtx()
	} else {
		txCtx = stores.DBTxStore.GetCtx(r.txId)
	}

	buff, err := backend.HandleQuery(r.query, txCtx)
	if err != nil {
		return &DBResponse{
			Body:   strings.NewReader(err.Error()),
			Status: InvalidQuery,
		}
	}

	return &DBResponse{
		Body:   buff.Reader(),
		Status: OK,
	}
}

type responseStatus = uint8

const (
	OK responseStatus = iota
	InvalidQuery
	AccessDenied
	ServerError
)
