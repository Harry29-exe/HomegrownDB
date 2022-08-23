package data

import (
	"HomegrownDB/dbsystem/schema/table"
)

type RowBuffer interface {
	GetRowSlot(dataContentLen int) []byte
	Free()
	Fields() uint16
	Tables() []table.Definition
}

type FieldPtr = uint16

const FieldPtrSize = 2

func (i *BaseRowBuffer) Tables() []table.Definition {
	return i.tables
}
