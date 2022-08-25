package querry

import "HomegrownDB/dbsystem/tx"

type DBRequest struct {
	txId  tx.Id
	query string
	user  any //not in use for now
}

func (r *DBRequest) Handle() {

}
