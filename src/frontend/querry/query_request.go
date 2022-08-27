package querry

import (
	"HomegrownDB/backend/executor/qrow"
	"HomegrownDB/dbsystem/tx"
)

type DBRequest struct {
	txId  tx.Id
	query string
	user  any //not in use for now
}

type DBResponse struct {
	rows qrow.RowBuffer
}

func (r *DBRequest) Handle() *DBResponse {

}
