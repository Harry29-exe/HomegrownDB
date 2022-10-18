// Package fsm - free space map is package holding
// implementation of database free space map
package fsm

import (
	"HomegrownDB/dbsystem/access/dbfs"
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/tx"
	"fmt"
	"os"
	"testing"
)

func CreateFreeSpaceMap(filepath string) *FreeSpaceMap {
	file, err := os.Create(filepath)
	if err != nil {
		panic(err.Error())
	}
	io, err := createNewIO(file)
	if err != nil {
		panic(err.Error())
	}

	return &FreeSpaceMap{io: io}
}

func LoadFreeSpaceMap(file *os.File) *FreeSpaceMap {
	io, err := loadIO(file)
	if err != nil {
		panic("")
	}

	return &FreeSpaceMap{io: io}
}

func CreateTestFreeSpaceMap(file dbfs.FileLike, _ *testing.T) *FreeSpaceMap {
	io, err := createNewIO(file)
	if err != nil {
		panic(err.Error())
	}
	return &FreeSpaceMap{io: io}
}

// FreeSpaceMap is data structure stores
// information about how much space each
// page has and helps find one with enough
// space to fit inserting tuple
type FreeSpaceMap struct {
	io *io
}

type page struct {
	header []byte
	data   []byte
}

// FindPage returns number of page with at least the amount of requested space,
// returns error if no page fulfill the requirements
func (f *FreeSpaceMap) FindPage(availableSpace uint16, ctx *tx.Ctx) (dbbs.PageId, error) {
	percentageSpace := uint8(availableSpace / availableSpaceDivider)
	if availableSpace%availableSpaceDivider > 0 {
		percentageSpace++
	}

	return f.findPage(percentageSpace, ctx)
}

// UpdatePage updates page free space which is set to availableSpace parameter value
func (f *FreeSpaceMap) UpdatePage(availableSpace uint16, pageId dbbs.PageId) {
	lastLayerPageIndex := pageId / uint32(leafNodeCount)
	nodeIndex := pageId - lastLayerPageIndex*uint32(leafNodeCount) + uint32(nonLeafNodeCount)
	pageIndex := lastLayerPageIndex + uint32(leafNodeCount) + 1
	f.updatePage(uint8(availableSpace/availableSpaceDivider), pageIndex, uint16(nodeIndex))
}

func (f *FreeSpaceMap) findPage(space uint8, ctx *tx.Ctx) (dbbs.PageId, error) {
	buffer := make([]byte, pageSize)

	var err internalError
	pageIndex, nodeIndex := uint32(0), uint16(0)
	lastPageIndex, lastNodeIndex := uint32(0), uint16(0)
	leafNodeVal := uint8(0)
	for {
		currentPage := f.io.getRPage(pageIndex, buffer)
		lastNodeIndex = nodeIndex
		nodeIndex, err = f.findLeafNode(space, currentPage.data)
		f.io.releaseRPage(pageIndex)
		switch err {
		case none:
			break
		case notUpdated:
			if pageIndex == 0 {
				return 0, NoFreeSpace{}
			} else {
				f.updatePage(leafNodeVal, lastPageIndex, lastNodeIndex)
				return f.findPage(space, ctx)
			}
		case corrupted:
			//todo implement me
			panic("Not implemented")
		case noSpace:
			return 0, NoFreeSpace{}
		}

		if pageIndex > uint32(leafNodeCount) {
			break
		}
		leafNodeVal = currentPage.data[nodeIndex]
		lastPageIndex = pageIndex
		pageIndex = f.getFsmPageIndex(nodeIndex, pageIndex)
	}

	pageIndexInLayer := pageIndex - uint32(leafNodeCount+1)
	return pageIndexInLayer*uint32(leafNodeCount) + uint32(nodeIndex-nonLeafNodeCount),
		nil
}

func (f *FreeSpaceMap) findLeafNode(space uint8, pageData []byte) (uint16, internalError) {
	var nodeIndex uint16 = 0
	if pageData[0] < space {
		return 0, notUpdated
	}

	var nextNodeIndex uint16
	for nodeIndex < nonLeafNodeCount {
		nextNodeIndex = f.getLeftNodeIndex(nodeIndex)
		if nextNodeIndex < nodeCount && pageData[nextNodeIndex] >= space {
			nodeIndex = nextNodeIndex
			continue
		}

		nextNodeIndex = f.getRightNodeIndex(nodeIndex)
		if nextNodeIndex < nodeCount && pageData[nextNodeIndex] >= space {
			nodeIndex = nextNodeIndex
			continue
		}

		return nodeIndex, corrupted
	}
	return nodeIndex, none
}

func (f *FreeSpaceMap) updatePage(space uint8, pageIndex uint32, nodeIndex uint16) {
	buffer := make([]byte, pageSize)

	p := f.io.getWPage(pageIndex, buffer)
	if p.data[nodeIndex] == space {
		f.io.releaseWPage(pageIndex)
		return
	}

	p.data[nodeIndex] = space
	for nodeIndex != 0 {
		parentIndex := f.getParentIndex(nodeIndex)
		left, right := f.getLeftNodeIndex(parentIndex), f.getRightNodeIndex(parentIndex)
		newValue := max(p.data[left], p.data[right])

		nodeIndex = parentIndex
		if p.data[nodeIndex] == newValue {
			break
		}
		p.data[nodeIndex] = newValue
	}

	f.io.flushPage(pageIndex, buffer)
	f.io.releaseWPage(pageIndex)

	if pageIndex != 0 {
		parentPageIndex, parentNodeIndex := f.getFsmParentPageIndex(pageIndex)
		f.updatePage(p.data[0], parentPageIndex, parentNodeIndex)
	}
}

func (f *FreeSpaceMap) getParentIndex(childNodeIndex uint16) uint16 {
	return (childNodeIndex - 1) / 2
}

func (f *FreeSpaceMap) getLeftNodeIndex(parentNodeIndex uint16) uint16 {
	return parentNodeIndex*2 + 1
}

func (f *FreeSpaceMap) getRightNodeIndex(parentNodeIndex uint16) uint16 {
	return parentNodeIndex*2 + 2
}

func (f *FreeSpaceMap) getFsmPageIndex(nodeIndex uint16, pageIndex uint32) uint32 {
	inLayerNodeIndex := uint32(nodeIndex - nonLeafNodeCount)
	if pageIndex == 0 {
		return inLayerNodeIndex + 1
	} else if pageIndex < uint32(leafNodeCount+1) {
		return inLayerNodeIndex +
			uint32(leafNodeCount)*(pageIndex-1) +
			uint32(leafNodeCount) + 1
	} else {
		panic(fmt.Sprintf("not expected that hight pageIndex = %d "+
			"(expected that fsm has 3 layers: 0, 1 and 2 so last valid pageIndex = leaftNodePerPage)",
			pageIndex),
		)
	}
}

func (f *FreeSpaceMap) getFsmParentPageIndex(pageIndex uint32) (parentPageIndex uint32, nodeIndex uint16) {
	if pageIndex > uint32(leafNodeCount) { // layer 2
		inLayerPageIndex := pageIndex - (uint32(leafNodeCount) + 1)

		parentPageIndex = inLayerPageIndex/uint32(leafNodeCount) + 1
		nodeIndex = nonLeafNodeCount + uint16(inLayerPageIndex) - uint16(parentPageIndex-1)*leafNodeCount
	} else if pageIndex > 0 { // layer 1
		parentPageIndex = 0
		nodeIndex = nonLeafNodeCount + uint16(pageIndex) - 1
	} else {
		panic("layer 0 can not has parent page")
	}
	return
}

func (f *FreeSpaceMap) leafNodeToPageId(nodeIndex uint16, pageIndex uint16, pageLayer uint16) uint16 {
	return (nodeIndex - (leafNodeCount - 1)) + pageLayer*leafNodeCount + pageIndex
}

func max(v1, v2 uint8) uint8 {
	if v1 > v2 {
		return v1
	}
	return v2
}
