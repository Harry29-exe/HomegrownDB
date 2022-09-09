package anode

import (
	"HomegrownDB/dbsystem/schema/table"
)

type QtableId = uint16

type Table struct {
	TableId  table.Id
	QtableId QtableId
	Alias    string
}
