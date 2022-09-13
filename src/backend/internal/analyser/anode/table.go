package anode

import (
	"HomegrownDB/dbsystem/schema/table"
)

type QtableId = uint16

type Table struct {
	Table    table.Definition
	QtableId QtableId
	Alias    string
}
