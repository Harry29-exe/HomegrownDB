// Package fsm - free space map is package holding
// implementation of database free space map
package fsm

import (
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
	"fmt"
)

// FreeSpaceMap is data structure stores
// information about how much space each
// page has and helps find one with enough
// space to fit inserting tuple
type FreeSpaceMap struct {
	io io

	table table.Definition
}

// FindPage returns number of page with at least the amount of requested space,
// returns error if no page fulfill the requirements
func (f *FreeSpaceMap) FindPage(availableSpace uint16, ctx *tx.Ctx) (dbbs.PageId, error) {
	//todo implement me
	panic("Not implemented")
}

// UpdatePage updates page free space which is set to availableSpace parameter value
func (f *FreeSpaceMap) UpdatePage(availableSpace uint16, pageId dbbs.PageId) {
	//todo implement me
	panic("Not implemented")
}

func (f *FreeSpaceMap) getInPageLeftChildIndex(inPageIndex uint16) uint16 {
	return inPageIndex*2 + 1
}

func (f *FreeSpaceMap) getInPageRightChildIndex(inPageIndex uint16) uint16 {
	return inPageIndex*2 + 2
}

func (f *FreeSpaceMap) getPageIndex(inPageNodeIndex uint16, pageIndex uint32) uint32 {
	inLayerNodeIndex := uint32(inPageNodeIndex - nonLeaftNodeCount)
	if pageIndex == 0 {
		return inLayerNodeIndex + 1
	} else if pageIndex < uint32(leafNodePerPage+1) {
		return inLayerNodeIndex +
			uint32(leafNodePerPage)*(pageIndex-1) +
			uint32(leafNodePerPage) + 1
	} else {
		panic(fmt.Sprintf("not expected that hight pageIndex = %d "+
			"(expected that fsm has 3 layers: 0, 1 and 2 so last valid pageIndex = leaftNodePerPage)",
			pageIndex),
		)
	}
}

func (f *FreeSpaceMap) leafNodeToPageId(inPageIndex uint16, pageIndex uint16, pageLayer uint16) uint16 {
	return (inPageIndex - (leafNodePerPage - 1)) + pageLayer*leafNodePerPage + pageIndex
}
