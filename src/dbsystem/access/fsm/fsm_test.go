package fsm

import (
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestFreeSpaceMap_getPageIndex(t *testing.T) {
	fsm := &FreeSpaceMap{}

	firstLeafNodeIndex := nonLeaftNodeCount
	secondLeafNodeIndex := nonLeaftNodeCount + 1
	middleLeafNodeIndex := nonLeaftNodeCount + leafNodePerPage/2
	penultimateLeafNodeIndex := nonLeaftNodeCount + leafNodePerPage - 2
	lastLeafNodePerIndex := nonLeaftNodeCount + leafNodePerPage - 1

	var pageInLayer uint16
	var layer uint32
	var newPagePreCalculatedOffset uint32

	testFsmGetPageIndex := func(nodeIndex uint16, layer uint32, pageInLayer uint16, offset uint32) {
		var pageIndex uint32
		if layer == 0 {
			pageIndex = uint32(pageInLayer)
		} else {
			pageIndex = uint32(pageInLayer) + 1
		}
		newPageIndex := fsm.getPageIndex(nodeIndex, pageIndex)
		assert.Eq(newPageIndex, offset+uint32(nodeIndex-nonLeaftNodeCount), t)
	}

	// page 0 layer 0
	pageInLayer = 0
	layer = 0
	newPagePreCalculatedOffset = 1

	testFsmGetPageIndex(firstLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(secondLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(middleLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(penultimateLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(lastLeafNodePerIndex, layer, pageInLayer, newPagePreCalculatedOffset)

	// page 0 layer 1
	pageInLayer = 0
	layer = 1
	newPagePreCalculatedOffset = uint32(1 + leafNodePerPage)

	testFsmGetPageIndex(firstLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(secondLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(middleLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(penultimateLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(lastLeafNodePerIndex, layer, pageInLayer, newPagePreCalculatedOffset)

	// page 1 layer 1
	pageInLayer = 1
	layer = 1
	newPagePreCalculatedOffset = uint32(1+leafNodePerPage) + uint32(leafNodePerPage)

	testFsmGetPageIndex(firstLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(secondLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(middleLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(penultimateLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(lastLeafNodePerIndex, layer, pageInLayer, newPagePreCalculatedOffset)

	// page middle layer 1
	pageInLayer = leafNodePerPage / 2
	layer = 1
	newPagePreCalculatedOffset =
		uint32(1+leafNodePerPage) + // layer 2 offset
			uint32(leafNodePerPage)*uint32(pageInLayer) // offset within layer 2

	testFsmGetPageIndex(firstLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(secondLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(middleLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(penultimateLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(lastLeafNodePerIndex, layer, pageInLayer, newPagePreCalculatedOffset)

	// page last, layer 1
	pageInLayer = leafNodePerPage - 1
	layer = 1
	newPagePreCalculatedOffset =
		uint32(1+leafNodePerPage) + // layer 2 offset
			uint32(leafNodePerPage)*uint32(pageInLayer) // offset within layer 2

	testFsmGetPageIndex(firstLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(secondLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(middleLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(penultimateLeafNodeIndex, layer, pageInLayer, newPagePreCalculatedOffset)
	testFsmGetPageIndex(lastLeafNodePerIndex, layer, pageInLayer, newPagePreCalculatedOffset)
}
