package exenode

import (
	"HomegrownDB/backend/executor/exenode/internal/data"
	"HomegrownDB/backend/planer/plan"
)

type ExeNode interface {
	// HasNext returns whether node can return more values
	HasNext() bool
	// Next returns next row
	Next() data.Row
	// NextBatch returns amount of rows that node can read in most effective way
	NextBatch() []data.Row
	// All returns all remaining rows
	All() []data.Row
}

type PlanTableId = plan.TableId
