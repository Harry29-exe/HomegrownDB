package qctx

import (
	"HomegrownDB/dbsystem/schema/column"
	"math"
)

type QTableId = uint16

var (
	InvalidQTableId QTableId = math.MaxUint16
	MaxQTableId     QTableId = math.MaxUint16 - 1
)

type QColumnId struct {
	QTableId QTableId
	ColOrder column.Order
}
