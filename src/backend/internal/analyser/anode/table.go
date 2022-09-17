package anode

import (
	"HomegrownDB/dbsystem/schema/table"
)

type QTableId = uint16

type Table struct {
	Table    table.Definition
	QtableId QTableId
	Alias    string
}
