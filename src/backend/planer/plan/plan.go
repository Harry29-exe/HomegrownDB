package plan

import "HomegrownDB/dbsystem/schema/table"

type Plan struct {
	RootNode Node

	Tables []Table
}

type Table struct {
	TableId     table.Id
	PlanTableId TableId
}

type TableId = uint8
