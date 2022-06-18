package node

import (
	"HomegrownDB/dbsystem/bstructs"
)

type Node interface {
	HasNext() bool
	Next() bstructs.Tuple
}
