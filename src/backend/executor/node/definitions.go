package node

import (
	"HomegrownDB/dbsystem/page/tuple"
)

type Node interface {
	HasNext() bool
	Next() tuple.Tuple
}
