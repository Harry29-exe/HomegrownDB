package page

import (
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/io/bparse"
)

type Tuple struct {
	data []byte
	//CreatedByTx      uint32
	//ModifiedByTx     uint32
	//TxCommandCounter uint32
	//Id               Id
	//
	//nullBitmap uint16
	//
	//columns []Column
}

func (t Tuple) CreatedByTx() (id tx.Id) {
	return bparse.Parse.UInt4(t.data)
}
