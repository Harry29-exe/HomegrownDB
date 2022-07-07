package node

import (
	"HomegrownDB/dbsystem/bdata"
)

type Node interface {
	HasNext() bool
	Next() bdata.Tuple
}
