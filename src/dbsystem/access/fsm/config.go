package fsm

import "HomegrownDB/dbsystem/dbbs"

const (
	pageLayers = 3
	headerSize = 8
)

var (
	pageSize     = dbbs.PageSize
	layersInPage uint16
)

func init() {
	layersInPage = (pageSize - headerSize) / 2
}
