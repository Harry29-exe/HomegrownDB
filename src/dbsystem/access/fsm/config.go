package fsm

import "HomegrownDB/dbsystem/dbbs"

const (
	pageLayers = 3
	headerSize = 8
)

var (
	pageSize          = dbbs.PageSize
	leafNodePerPage   uint16
	nonLeaftNodeCount uint16
	pageLayerOffsets  []uint32
)

func init() {
	leafNodePerPage = (pageSize - headerSize) / 2
	nonLeaftNodeCount = pageSize - leafNodePerPage
	pageLayerOffsets = []uint32{
		0,                           // root has no offset
		1,                           // root layer has one page
		uint32(leafNodePerPage) + 1, // 1'st layer has leafNodePerPage pages and root has one page
	}
}
