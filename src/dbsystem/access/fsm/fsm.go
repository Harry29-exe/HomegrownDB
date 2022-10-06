// Package fsm - free space map is package holding
// implementation of database free space map
package fsm

import (
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
)

func CreateNewFSM(table table.Definition) *FreeSpaceMap {

}

func InitFSM(table table.Definition) *FreeSpaceMap {

}

// FreeSpaceMap is data structure stores
// information about how much space each
// page has and helps find one with enough
// space to fit inserting tuple
type FreeSpaceMap struct {
	io IO

	table  table.Definition
	layers uint8
}

// FindPage returns number of page with at least the amount of requested space,
// returns error if no page fulfill the requirements
func (f *FreeSpaceMap) FindPage(availableSpace uint16, ctx *tx.Ctx) (dbbs.PageId, error) {
	//todo implement me
	panic("Not implemented")
}

// UpdatePage updates page free space which is set to availableSpace parameter value
func (f *FreeSpaceMap) UpdatePage(availableSpace uint16, pageId dbbs.PageId) {
	//todo implement me
	panic("Not implemented")
}
