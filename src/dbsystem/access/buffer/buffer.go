package buffer

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/storage/fsm/fsmpage"
	"HomegrownDB/dbsystem/storage/page"
)

type SharedBuffer interface {
	TableBuffer
	FsmBuffer

	WPageRelease(tag page.PageTag)
	RPageRelease(tag page.PageTag)
	FlushAll() error
}

type TableBuffer interface {
	RTablePage(table tabdef.RDefinition, pageId page.Id) (page.RPage, error)
	WTablePage(table tabdef.RDefinition, pageId page.Id) (page.WPage, error)
}

type FsmBuffer interface {
	RFsmPage(ownerID hglib.OID, pageId page.Id) (fsmpage.Page, error)
	WFsmPage(ownerID hglib.OID, pageId page.Id) (fsmpage.Page, error)
}

const NewPage page.Id = page.InvalidId

type TableSrc interface {
	Table(id tabdef.Id) tabdef.RDefinition
}

type slotIndex = uint

var pageSize = int64(page.Size)
