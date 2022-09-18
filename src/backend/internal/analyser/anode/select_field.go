package anode

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
)

type SelectFields struct {
	Fields []SelectField
}

type QFieldId = uint16

type SelectField struct {
	Table      table.Definition
	Column     column.Definition
	FieldAlias string
}
