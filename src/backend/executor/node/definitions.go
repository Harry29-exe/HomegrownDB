package node

import (
	"HomegrownDB/dbsystem/page"
)

type Node interface {
	HasNext() bool
	Next() page.Tuple
}
