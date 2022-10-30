// Package fsm - free space map is package holding
// implementation of database free space map
package fsm

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

func CreateFreeSpaceMap(rel relation.Relation, buff buffer.SharedBuffer) (*FreeSpaceMap, error) {
	fsm := &FreeSpaceMap{rel: rel, buff: buff}
	if err := initNewFsmIO(fsm); err != nil {
		return fsm, err
	}

	return fsm, nil
}

func LoadFreeSpaceMap(rel relation.Relation, buff buffer.SharedBuffer) (*FreeSpaceMap, error) {
	return &FreeSpaceMap{rel: rel, buff: buff}, nil
}

func initNewFsmIO(fsm *FreeSpaceMap) error {
	for i := 0; i < int(leafNodeCount+1); i++ {
		tag := buffer.NewPageTag(page.Id(i), fsm.rel)
		_, err := fsm.buff.WGenericPage(tag, fsm.rel)
		if err != nil {
			return err
		}
		fsm.buff.ReleaseWPage(tag)
	}
	return nil
}

// FreeSpaceMap is data structure stores
// information about how much space each
// page has and helps find one with enough
// space to fit inserting tuple
type FreeSpaceMap struct {
	rel  relation.Relation
	buff buffer.SharedBuffer
}

// FindPage returns number of page with at least the amount of requested space,
// returns error if no page fulfill the requirements
func (f *FreeSpaceMap) FindPage(availableSpace uint16, ctx *tx.Ctx) (page.Id, error) {
	percentageSpace := uint8(availableSpace / availableSpaceDivider)
	if availableSpace%availableSpaceDivider > 0 {
		percentageSpace++
	}

	return f.findPage(percentageSpace, ctx)
}

// UpdatePage updates page free space which is set to availableSpace parameter value
func (f *FreeSpaceMap) UpdatePage(availableSpace uint16, pageId page.Id) error {
	lastLayerPageIndex := pageId / uint32(leafNodeCount)
	nodeIndex := pageId - lastLayerPageIndex*uint32(leafNodeCount) + uint32(nonLeafNodeCount)
	pageIndex := lastLayerPageIndex + uint32(leafNodeCount) + 1
	return f.updatePages(uint8(availableSpace/availableSpaceDivider), pageIndex, uint16(nodeIndex))
}
