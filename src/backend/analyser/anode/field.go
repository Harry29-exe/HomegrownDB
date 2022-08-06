package anode

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
)

type Fields struct {
	Fields []Field
}

type Field struct {
	Table     table.Definition
	Column    column.Definition
	FieldName string
}
