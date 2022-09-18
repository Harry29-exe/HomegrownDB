package anode

import "HomegrownDB/dbsystem/schema/column"

type Insert struct {
	Table   Table
	Columns []column.OrderId
	Values  Values
}
