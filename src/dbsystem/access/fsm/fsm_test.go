package fsm

import (
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestFreeSpaceMap_getPageIndex(t *testing.T) {
	fsm := &FreeSpaceMap{}

	pageIndex := fsm.getPageIndex(nonLeaftNodeCount+1, 1)
	assert.Eq(pageIndex, uint32(leafNodePerPage+1+1), t)

	pageIndex = fsm.getPageIndex(nonLeaftNodeCount+1, 0)
	assert.Eq(pageIndex, 2, t)

}
