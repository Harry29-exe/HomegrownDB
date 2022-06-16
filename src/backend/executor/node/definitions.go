package node

import (
	"HomegrownDB/dbsystem/buffer"
)

type Node interface {
	HasNext() bool
	Next() buffer.Tuple
}
