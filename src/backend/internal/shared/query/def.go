package query

import "HomegrownDB/dbsystem/schema/table"

type QTableId = int

type QTable struct {
	Id    QTableId
	Table table.Definition
}
