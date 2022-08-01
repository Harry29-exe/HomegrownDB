package anode

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
)

type Fields struct {
	fields []Field
}

type Field struct {
	table  table.Definition
	column column.Definition
}
