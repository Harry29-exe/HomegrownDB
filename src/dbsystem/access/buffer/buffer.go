package buffer

import (
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/fsm/fsmpage"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/tpage"
)

type SharedBuffer interface {
	TableBuffer
	FsmBuffer

	WPageRelease(tag page.PageTag)
	RPageRelease(tag page.PageTag)
}

type TableBuffer interface {
	RTablePage(table table.RDefinition, pageId page.Id) (tpage.RPage, error)
	WTablePage(table table.RDefinition, pageId page.Id) (tpage.WPage, error)
}

type FsmBuffer interface {
	RFsmPage(ownerID dbobj.OID, pageId page.Id) (fsmpage.Page, error)
	WFsmPage(ownerID dbobj.OID, pageId page.Id) (fsmpage.Page, error)
}

const NewPage page.Id = page.InvalidId

type TableSrc interface {
	Table(id table.Id) table.RDefinition
}

type slotIndex = uint

var pageSize = int64(page.Size)
