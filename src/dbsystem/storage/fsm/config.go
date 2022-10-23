package fsm

import (
	"HomegrownDB/dbsystem/storage/page"
)

const (
	pageLayers = 3
	headerSize = 8
)

var (
	availableSpaceDivider uint16
	pageSize              = page.Size
	leafNodeCount         uint16
	nonLeafNodeCount      uint16
	nodeCount             uint16
)

func init() {
	nonLeafNodeCount = pageSize/2 - 1
	leafNodeCount = (pageSize - headerSize) - nonLeafNodeCount
	nodeCount = leafNodeCount + nonLeafNodeCount
	availableSpaceDivider = page.Size / 256
}
