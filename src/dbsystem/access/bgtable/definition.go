package bgtable

import (
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/schema/column"
	"encoding/binary"
)

type BgId = uint64

func CreateBgId(tid bdata.TID, columnOrder column.Order) BgId {
	id := make([]byte, 8)
	binary.BigEndian.PutUint32(id, tid.PageId)
	binary.BigEndian.PutUint16(id[4:], tid.TupleIndex)
	binary.BigEndian.PutUint16(id[6:], columnOrder)

	return binary.BigEndian.Uint64(id)
}
