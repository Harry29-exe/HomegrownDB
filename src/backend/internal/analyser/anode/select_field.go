package anode

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/tx"
)

type SelectFields = []SelectField

type QFieldId = uint16

type SelectField struct {
	Table  tx.QTableId
	Column column.Order
}
