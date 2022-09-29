package plan

import (
	"HomegrownDB/dbsystem/schema/column"
)

type FieldId struct {
	PlanTableId TableId
	ColumnOrder column.Order
}
