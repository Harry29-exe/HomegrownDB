package node

import (
	"HomegrownDB/sql/pager/tuple"
)

type Node interface {
	HasNext() bool
	Next() tuple.Tuple
}
