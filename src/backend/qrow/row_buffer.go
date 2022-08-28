package qrow

import (
	"HomegrownDB/dbsystem/schema/table"
	"io"
)

type RowBuffer interface {
	GetRowSlot(dataContentLen int) []byte
	Free()
	Fields() uint16
	Tables() []table.Definition
	Reader() io.Reader
}

type FieldPtr = uint16

const FieldPtrSize = 2
