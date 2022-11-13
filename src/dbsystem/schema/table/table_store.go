package table

import "math"

type Store interface {
	WDefsStore
}

type RDefsStore interface {
	// FindTable searches for table with given name
	// if it exists it returns table id otherwise it
	// returns InvalidTableId
	FindTable(name string) Id

	AccessTable(id Id, lockMode tableLockMode) Definition
}

type WDefsStore interface {
	RDefsStore
	AddTable(table Definition) error
	RemoveTable(id Id) error
}

var (
	InvalidTableId Id = math.MaxUint16
	MaxTableId     Id = math.MaxUint16 - 1
)

type tableLockMode = uint8

const (
	RLockMode tableLockMode = iota
	WLockMode
)
