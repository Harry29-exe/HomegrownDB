package exenode

import (
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/backend/internal/shared/query"
)

type ExeNode interface {
	// SetSource sets from node will receive incoming nodes
	// it's usually one node but in some cases it can be multiple nodes (e.g. Join)
	// and sometimes it will be nil (e.g. SeqScan)
	SetSource(source []ExeNode)
	// HasNext returns whether node can return more values
	HasNext() bool
	// Next returns next row
	Next() query.QRow
	// NextBatch returns amount of rows that node can read in most effective way
	NextBatch() []query.QRow
	// All returns all remaining rows
	All() []query.QRow
	// Free send signal to node that it won't be used and should release
	// all its resources
	Free()
}

type PlanTableId = plan.TableId

type InitOptions struct {
	StoreAllRows bool
}
