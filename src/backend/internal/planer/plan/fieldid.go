package plan

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
)

type FieldId struct {
	PlanTableId table.Id
	ColumnOrder column.Order
}
