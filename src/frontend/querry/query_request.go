package querry

import (
	"HomegrownDB/dbsystem/tx"
	"io"
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
	//var txCtx *tx.Ctx
	//if r.txId == 0 {
	//	txCtx = tx.DBTxStore.NewCtx()
	//} else {
	//	txCtx = tx.DBTxStore.GetCtx(r.txId)
	//}

	//todo implement me
	panic("Not implemented")
	//buff, err := backend.HandleQuery(r.query, txCtx)
	//if err != nil {
	//	return &DBResponse{
	//		Body:   strings.NewReader(err.Error()),
	//		TxStatus: InvalidQuery,
	//	}
	//}
	//
	//return &DBResponse{
	//	Body:   buff.Reader(),
	//	TxStatus: OK,
	//}
}

type responseStatus = uint8

const (
	OK responseStatus = iota
	InvalidQuery
	AccessDenied
	ServerError
)
