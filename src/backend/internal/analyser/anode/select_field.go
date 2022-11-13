package anode

import (
	"HomegrownDB/backend/internal/shared/qctx"
	"HomegrownDB/dbsystem/schema/column"
)

type SelectFields = []SelectField

type QFieldId = uint16

type SelectField struct {
	Table  qctx.QTableId
	Column column.Order
}
