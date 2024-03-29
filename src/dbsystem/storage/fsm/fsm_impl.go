package fsm

import (
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"fmt"
)

func (f *FSM) findPage(space uint8, ctx tx.Tx) (page.Id, error) {
	var internalErr internalError
	pageIndex, nodeIndex := uint32(0), uint16(0)
	lastPageIndex, lastNodeIndex := uint32(0), uint16(0)
	newLeafNodeVal, leafNodeVal := uint8(0), uint8(0)

	for {
		pageTag := page.NewPageTag(pageIndex, f.fsmOID)
		rPage, err := f.buff.RFsmPage(f.fsmOID, pageIndex)
		if err != nil {
			return 0, err
		}
		pageData := rPage.Data()

		lastNodeIndex = nodeIndex
		nodeIndex, internalErr = f.findLeafNode(space, pageData)
		newLeafNodeVal = pageData[nodeIndex]
		f.buff.RPageRelease(pageTag)

		if internalErr == corrupted {
			//todo implement me
			panic("Not implemented")
		} else if internalErr == noSpace {
			if pageIndex == 0 {
				return page.InvalidId, nil
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
		leafNodeVal = newLeafNodeVal
		lastPageIndex = pageIndex
		pageIndex = f.getFsmPageIndex(nodeIndex, pageIndex)
	}

	return f.calcPageId(pageIndex, nodeIndex), nil
}

func (f *FSM) calcPageId(pageIndex uint32, nodeIndex uint16) page.Id {
	pageIndexInLayer := pageIndex - uint32(leafNodeCount+1)
	return pageIndexInLayer*uint32(leafNodeCount) + uint32(nodeIndex-nonLeafNodeCount)
}

func (f *FSM) findLeafNode(space uint8, pageData []byte) (uint16, internalError) {
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

func (f *FSM) updatePages(space uint8, pageIndex uint32, nodeIndex uint16) error {
	tag := page.NewPageTag(pageIndex, f.fsmOID)

	wPage, err := f.buff.WFsmPage(f.fsmOID, pageIndex)
	if err != nil {
		return err
	}
	pageData := wPage.Data()
	if pageData[nodeIndex] == space {
		f.buff.WPageRelease(tag)
		return nil
	}

	f.updatePage(space, pageData, nodeIndex)

	f.buff.WPageRelease(tag)

	if pageIndex != 0 {
		parentPageIndex, parentNodeIndex := f.getFsmParentPageIndex(pageIndex)
		return f.updatePages(pageData[0], parentPageIndex, parentNodeIndex)
	}
	return nil
}

func (f *FSM) updatePage(space uint8, pageData []byte, nodeIndex uint16) {
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

func (f *FSM) getParentIndex(childNodeIndex uint16) uint16 {
	return (childNodeIndex - 1) / 2
}

func (f *FSM) getLeftNodeIndex(parentNodeIndex uint16) uint16 {
	return parentNodeIndex*2 + 1
}

func (f *FSM) getRightNodeIndex(parentNodeIndex uint16) uint16 {
	return parentNodeIndex*2 + 2
}

func (f *FSM) getFsmPageIndex(nodeIndex uint16, pageIndex uint32) uint32 {
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

func (f *FSM) getFsmParentPageIndex(pageIndex uint32) (parentPageIndex uint32, nodeIndex uint16) {
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

func (f *FSM) leafNodeToPageId(nodeIndex uint16, pageIndex uint16, pageLayer uint16) uint16 {
	return (nodeIndex - (leafNodeCount - 1)) + pageLayer*leafNodeCount + pageIndex
}

func max(v1, v2 uint8) uint8 {
	if v1 > v2 {
		return v1
	}
	return v2
}
