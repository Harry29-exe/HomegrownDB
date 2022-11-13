package exenode

import (
	"HomegrownDB/backend/internal/shared/query"
)

type ExeNode interface {
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

type InitOptions struct {
	StoreAllRows bool
}
