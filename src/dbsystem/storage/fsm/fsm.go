// Package fsm - free space map is package holding
// implementation of database free space map
package fsm

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
)

// CreateFreeSpaceMap creates new DBObject directory and initializes its data file
func CreateFreeSpaceMap(
	fsmOID dbobj.OID,
	fs dbfs.FS,
) error {
	err := fs.InitNewRelationDir(fsmOID)
	if err != nil {
		return err
	}
	file, err := fs.OpenRelationDataFile(fsmOID)
	if err != nil {
		return err
	}
	io, err := pageio.NewPageIO(file)
	if err != nil {
		return err
	}

	pagesToCreate := leafNodeCount + 1
	lastPageIndex := pagesToCreate - 1
	if io.FlushPage(page.Id(lastPageIndex), make([]byte, page.Size)) != nil {
		return err
	}
	return nil
}

func NewFSM(fsmOID dbobj.OID, buff buffer.SharedBuffer) *FSM {
	return &FSM{
		fsmOID: fsmOID,
		buff:   buff,
	}
}

func initNewFsmIO(fsm *FSM) error {
	for i := 0; i < int(leafNodeCount+1); i++ {
		tag := page.NewPageTag(page.Id(i), fsm.fsmOID)
		_, err := fsm.buff.WFsmPage(fsm.fsmOID, page.Id(i))
		if err != nil {
			return err
		}
		fsm.buff.WPageRelease(tag)
	}
	return nil
}

// FSM is data structure stores
// information about how much space each
// page has and helps find one with enough
// space to fit inserting tuple
type FSM struct {
	fsmOID dbobj.OID
	buff   buffer.SharedBuffer
}

// FindPage returns number of page with at least the amount of requested space,
// if no page fulfill the requirements returns page.InvalidId
func (f *FSM) FindPage(availableSpace uint16, tx tx.Tx) (page.Id, error) {
	percentageSpace := uint8(availableSpace / availableSpaceDivider)
	if availableSpace%availableSpaceDivider > 0 {
		percentageSpace++
	}

	return f.findPage(percentageSpace, tx)
}

// UpdatePage updates page free space which is set to availableSpace parameter value
func (f *FSM) UpdatePage(availableSpace uint16, pageId page.Id) error {
	lastLayerPageIndex := pageId / uint32(leafNodeCount)
	nodeIndex := pageId - lastLayerPageIndex*uint32(leafNodeCount) + uint32(nonLeafNodeCount)
	pageIndex := lastLayerPageIndex + uint32(leafNodeCount) + 1
	return f.updatePages(uint8(availableSpace/availableSpaceDivider), pageIndex, uint16(nodeIndex))
}
