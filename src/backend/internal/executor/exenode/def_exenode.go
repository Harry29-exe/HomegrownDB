package exenode

import (
	"HomegrownDB/backend/internal/planer/plan"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/dbbs"
	"HomegrownDB/dbsystem/tx"
)

type ExeNode interface {
	// SetSource sets from node will receive incoming nodes
	// it's usually one node but in some cases it can be multiple nodes (e.g. Join)
	// and sometimes it will be nil (e.g. SeqScan)
	SetSource(source []ExeNode)
	// HasNext returns whether node can return more values
	HasNext() bool
	// Next returns next row
	Next() dbbs.QRow
	// NextBatch returns amount of rows that node can read in most effective way
	NextBatch() []dbbs.QRow
	// All returns all remaining rows
	All() []dbbs.QRow
	// Free send signal to node that it won't be used and should release
	// all its resources
	Free()
}

type PlanTableId = plan.TableId

type InitOptions struct {
	StoreAllRows bool
}

type Builder interface {
	Build(node plan.Node, ctx BuildCtx) ExeNode
}

type BuildCtx = *buildCtx

type buildCtx struct {
	tx   *tx.Ctx
	buff buffer.SharedBuffer
}
