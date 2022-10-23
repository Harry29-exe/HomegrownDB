// Package fsm - free space map is package holding
// implementation of database free space map
package fsm

import (
	"HomegrownDB/dbsystem/access/dbbs"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
	"os"
	"testing"
)

func CreateFreeSpaceMap(filepath string) (*FreeSpaceMap, error) {
	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}
	pageIO, err := pageio.NewRWPageIO(file)
	if err != nil {
		return nil, err
	}

	if err = initNewFsmIO(pageIO); err != nil {
		return nil, err
	}

	return &FreeSpaceMap{io: pageIO}, nil
}

func LoadFreeSpaceMap(file dbfs.FileLike) (*FreeSpaceMap, error) {
	pageIO, err := pageio.LoadRWPageIO(file)
	if err != nil {
		return nil, err
	}

	return &FreeSpaceMap{io: pageIO}, nil
}

func CreateTestFreeSpaceMap(file dbfs.FileLike, _ *testing.T) (*FreeSpaceMap, error) {
	pageIO, err := pageio.NewRWPageIO(file)
	if err != nil {
		return nil, err
	}
	if err = initNewFsmIO(pageIO); err != nil {
		return nil, err
	}

	return &FreeSpaceMap{io: pageIO}, nil
}

func initNewFsmIO(io pageio.ResourceLockIO) error {
	buff := make([]byte, pageSize)
	for i := 0; i < int(leafNodeCount+1); i++ {
		_, err := io.NewPage(buff)
		if err != nil {
			return err
		}
	}
	return nil
}

// FreeSpaceMap is data structure stores
// information about how much space each
// page has and helps find one with enough
// space to fit inserting tuple
type FreeSpaceMap struct {
	io pageio.ResourceLockIO
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
func (f *FreeSpaceMap) UpdatePage(availableSpace uint16, pageId dbbs.PageId) error {
	lastLayerPageIndex := pageId / uint32(leafNodeCount)
	nodeIndex := pageId - lastLayerPageIndex*uint32(leafNodeCount) + uint32(nonLeafNodeCount)
	pageIndex := lastLayerPageIndex + uint32(leafNodeCount) + 1
	return f.updatePages(uint8(availableSpace/availableSpaceDivider), pageIndex, uint16(nodeIndex))
}
