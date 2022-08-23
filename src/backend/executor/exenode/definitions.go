package exenode

import (
	"HomegrownDB/backend/executor/exenode/internal/data"
	"HomegrownDB/backend/planer/plan"
)

type ExeNode interface {
	// Init initialize node with dataHolder that will be used to create data.R
	Init(dataHolder data.RowBuffer)
	// HasNext returns whether node can return more values
	HasNext() bool
	// Next returns next row
	Next() data.Row
	// NextBatch returns amount of rows that node can read in most effective way
	NextBatch() []data.Row
	// All returns all remaining rows
	All() []data.Row
	// Free send signal to node that it won't be used and should release
	// all its resources
	Free()
}

type PlanTableId = plan.TableId
