// Package fsm - free space map is package holding
// implementation of database free space map
package fsm

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/access/buffer"
	relation "HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
)

func CreateFreeSpaceMap(fsmRelation relation.BaseRelation, parentRelationId relation.ID, buff buffer.SharedBuffer) (*FreeSpaceMap, error) {
	fsm := &FreeSpaceMap{
		BaseRelation:     fsmRelation,
		parentRelationId: parentRelationId,
		buff:             buff,
	}
	if err := initNewFsmIO(fsm); err != nil {
		return fsm, err
	}

	return fsm, nil
}

func LoadFreeSpaceMap(fsmRelation relation.BaseRelation, parentRelationId relation.ID, buff buffer.SharedBuffer) (*FreeSpaceMap, error) {
	return &FreeSpaceMap{
		BaseRelation:     fsmRelation,
		parentRelationId: parentRelationId,
		buff:             buff,
	}, nil
}

func initNewFsmIO(fsm *FreeSpaceMap) error {
	for i := 0; i < int(leafNodeCount+1); i++ {
		tag := pageio.NewPageTag(page.Id(i), fsm)
		_, err := fsm.buff.WFsmPage(fsm, page.Id(i))
		if err != nil {
			return err
		}
		fsm.buff.WPageRelease(tag)
	}
	return nil
}

func SerializeFSM(fsm *FreeSpaceMap, serializer *bparse.Serializer) {
	relation.SerializeBaseRelation(&fsm.BaseRelation, serializer)
	serializer.Uint32(uint32(fsm.parentRelationId))
}

func DeserializeFSM(buff buffer.SharedBuffer, deserializer *bparse.Deserializer) *FreeSpaceMap {
	return &FreeSpaceMap{
		BaseRelation:     relation.DeserializeBaseRelation(deserializer),
		parentRelationId: relation.ID(deserializer.Uint32()),
		buff:             buff,
	}
}

// FreeSpaceMap is data structure stores
// information about how much space each
// page has and helps find one with enough
// space to fit inserting tuple
type FreeSpaceMap struct {
	relation.BaseRelation
	parentRelationId relation.ID
	buff             buffer.SharedBuffer
}

// FindPage returns number of page with at least the amount of requested space,
// if no page fulfill the requirements returns page.InvalidId
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

var _ relation.Relation = &FreeSpaceMap{}
