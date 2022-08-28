package exenode

import (
	plan2 "HomegrownDB/backend/internal/planer/plan"
	qrow2 "HomegrownDB/backend/qrow"
)

type ExeNode interface {
	// SetSource sets from node will receive incoming nodes
	// it's usually one node but in some cases it can be multiple nodes (e.g. Join)
	// and sometimes it will be nil (e.g. SeqScan)
	SetSource(source []ExeNode)
	// Init initialize node with dataHolder that will be used to create data.R
	Init(options InitOptions) qrow2.RowBuffer
	// HasNext returns whether node can return more values
	HasNext() bool
	// Next returns next row
	Next() qrow2.Row
	// NextBatch returns amount of rows that node can read in most effective way
	NextBatch() []qrow2.Row
	// All returns all remaining rows
	All() []qrow2.Row
	// Free send signal to node that it won't be used and should release
	// all its resources
	Free()
}

type PlanTableId = plan2.TableId

type InitOptions struct {
	StoreAllRows bool
}

type Builder interface {
	Build(node plan2.Node) ExeNode
}
