package plan

import (
	"HomegrownDB/dbsystem/schema/table"
)

type SeqScan struct {
	Table      table.Definition
	Conditions Conditions
}
