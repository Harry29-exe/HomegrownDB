package fsm

import (
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"fmt"
)

func (f *FreeSpaceMap) findPage(space uint8, ctx *tx.Ctx) (page.Id, error) {
	buffer := make([]byte, pageSize)

	var internalErr internalError
	pageIndex, nodeIndex := uint32(0), uint16(0)
	lastPageIndex, lastNodeIndex := uint32(0), uint16(0)
	leafNodeVal := uint8(0)
	for {
		err := f.io.RPage(pageIndex, buffer)
		if err != nil {
			return 0, err
		}
		pageData := buffer[headerSize:]

		lastNodeIndex = nodeIndex
		nodeIndex, internalErr = f.findLeafNode(space, pageData)
		f.io.ReleaseRPage(pageIndex)

		switch internalErr {
		case none:
			break
		case corrupted:
			//todo implement me
			panic("Not implemented")
		case noSpace:
			if pageIndex == 0 {
				return 0, NoFreeSpace{}
			}
			err = f.updatePages(leafNodeVal, lastPageIndex, lastNodeIndex)
			if err != nil {
				return 0, err
			}
			return f.findPage(space, ctx)
		}

		if pageIndex > uint32(leafNodeCount) {
			break
		}
		leafNodeVal = pageData[nodeIndex]
		lastPageIndex = pageIndex
		pageIndex = f.getFsmPageIndex(nodeIndex, pageIndex)
	}

	return f.calcPageId(pageIndex, nodeIndex), nil
}

func (f *FreeSpaceMap) calcPageId(pageIndex uint32, nodeIndex uint16) page.Id {
	pageIndexInLayer := pageIndex - uint32(leafNodeCount+1)
	return pageIndexInLayer*uint32(leafNodeCount) + uint32(nodeIndex-nonLeafNodeCount)
}

func (f *FreeSpaceMap) findLeafNode(space uint8, pageData []byte) (uint16, internalError) {
	var nodeIndex uint16 = 0
	if pageData[0] < space {
		return 0, noSpace
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

func (f *FreeSpaceMap) updatePages(space uint8, pageIndex uint32, nodeIndex uint16) error {
	buffer := make([]byte, pageSize)

	err := f.getWritePage(pageIndex, buffer)
	if err != nil {
		return err
	}
	pageData := buffer[headerSize:]
	if pageData[nodeIndex] == space {
		f.io.ReleaseWPage(pageIndex)
		return nil
	}

	f.updatePage(space, pageData, nodeIndex)

	err = f.io.Flush(pageIndex, buffer)
	f.io.ReleaseWPage(pageIndex)
	if err != nil {
		return err
	}

	if pageIndex != 0 {
		parentPageIndex, parentNodeIndex := f.getFsmParentPageIndex(pageIndex)
		return f.updatePages(pageData[0], parentPageIndex, parentNodeIndex)
	}
	return nil
}

func (f *FreeSpaceMap) updatePage(space uint8, pageData []byte, nodeIndex uint16) {
	pageData[nodeIndex] = space
	for nodeIndex != 0 {
		parentIndex := f.getParentIndex(nodeIndex)
		left, right := f.getLeftNodeIndex(parentIndex), f.getRightNodeIndex(parentIndex)
		newValue := max(pageData[left], pageData[right])

		nodeIndex = parentIndex
		if pageData[nodeIndex] == newValue {
			break
		}
		pageData[nodeIndex] = newValue
	}
}

func (f *FreeSpaceMap) getWritePage(pageIndex uint32, buffer []byte) error {
	for f.io.PageCount() <= pageIndex {
		_, err := f.io.NewPage(buffer)
		if err != nil {
			return err
		}
	}

	return f.io.WPage(pageIndex, buffer)
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
