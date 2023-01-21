package fsm

import (
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestFreeSpaceMap_getPageIndex(t *testing.T) {
	fsm := &FSM{}

	firstLeafNodeIndex := nonLeafNodeCount
	secondLeafNodeIndex := nonLeafNodeCount + 1
	middleLeafNodeIndex := nonLeafNodeCount + leafNodeCount/2
	penultimateLeafNodeIndex := nonLeafNodeCount + leafNodeCount - 2
	lastLeafNodePerIndex := nonLeafNodeCount + leafNodeCount - 1

	var pageInLayer uint16
	var layer uint32
	var newPagePreCalculatedOffset uint32

	testFsmGetPageIndex := func(layer uint32, pageInLayer uint16, offset uint32) {
		var pageIndex uint32
		if layer == 0 {
			pageIndex = uint32(pageInLayer)
		} else {
			pageIndex = uint32(pageInLayer) + 1
		}

		newPageIndex := fsm.getFsmPageIndex(firstLeafNodeIndex, pageIndex)
		assert.Eq(newPageIndex, offset+uint32(firstLeafNodeIndex-nonLeafNodeCount), t)
		newPageIndex = fsm.getFsmPageIndex(secondLeafNodeIndex, pageIndex)
		assert.Eq(newPageIndex, offset+uint32(secondLeafNodeIndex-nonLeafNodeCount), t)
		newPageIndex = fsm.getFsmPageIndex(middleLeafNodeIndex, pageIndex)
		assert.Eq(newPageIndex, offset+uint32(middleLeafNodeIndex-nonLeafNodeCount), t)
		newPageIndex = fsm.getFsmPageIndex(penultimateLeafNodeIndex, pageIndex)
		assert.Eq(newPageIndex, offset+uint32(penultimateLeafNodeIndex-nonLeafNodeCount), t)
		newPageIndex = fsm.getFsmPageIndex(lastLeafNodePerIndex, pageIndex)
		assert.Eq(newPageIndex, offset+uint32(lastLeafNodePerIndex-nonLeafNodeCount), t)
	}

	// page 0 layer 0
	pageInLayer = 0
	layer = 0
	newPagePreCalculatedOffset = 1
	testFsmGetPageIndex(layer, pageInLayer, newPagePreCalculatedOffset)

	// page 0 layer 1
	pageInLayer = 0
	layer = 1
	newPagePreCalculatedOffset = uint32(1 + leafNodeCount)
	testFsmGetPageIndex(layer, pageInLayer, newPagePreCalculatedOffset)

	// page 1 layer 1
	pageInLayer = 1
	layer = 1
	newPagePreCalculatedOffset = uint32(1+leafNodeCount) + uint32(leafNodeCount)
	testFsmGetPageIndex(layer, pageInLayer, newPagePreCalculatedOffset)

	// page middle layer 1
	pageInLayer = leafNodeCount / 2
	layer = 1
	newPagePreCalculatedOffset =
		uint32(1+leafNodeCount) + // layer 2 offset
			uint32(leafNodeCount)*uint32(pageInLayer) // offset within layer 2
	testFsmGetPageIndex(layer, pageInLayer, newPagePreCalculatedOffset)

	// page last, layer 1
	pageInLayer = leafNodeCount - 1
	layer = 1
	newPagePreCalculatedOffset =
		uint32(1+leafNodeCount) + // layer 2 offset
			uint32(leafNodeCount)*uint32(pageInLayer) // offset within layer 2
	testFsmGetPageIndex(layer, pageInLayer, newPagePreCalculatedOffset)
}
