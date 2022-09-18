package anode

import (
	"HomegrownDB/dbsystem/schema/table"
)

// QTableId id of table in currently parsing query
type QTableId = uint16

type Table struct {
	Def      table.Definition
	QTableId QTableId
	Alias    string
}
