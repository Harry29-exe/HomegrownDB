// Package fsm - free space map is package holding
// implementation of database free space map
package fsm

import (
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
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
	pageOffset := uint32(1)
	pageIndexInLayer := pageIndex - 1
	pageLayer := 1 // starting from second layer
	for pageLayerOffsets[pageLayer] < pageIndex {
		pageOffset += pageLayerOffsets[pageLayer]
		pageIndexInLayer -= pageLayerOffsets[pageLayer]
	}

	return pageLayerOffsets[pageLayer+1] +
		pageIndexInLayer*uint32(leafNodePerPage) +
		uint32(inPageNodeIndex-nonLeaftNodeCount)
}

func (f *FreeSpaceMap) leafNodeToPageId(inPageIndex uint16, pageIndex uint16, pageLayer uint16) uint16 {
	return (inPageIndex - (leafNodePerPage - 1)) + pageLayer*leafNodePerPage + pageIndex
}
