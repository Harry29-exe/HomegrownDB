package anode

import "HomegrownDB/dbsystem/schema/table"

type Tables struct {
	Tables []Table
}

type Table struct {
	Table table.Definition
	Alias string
}
