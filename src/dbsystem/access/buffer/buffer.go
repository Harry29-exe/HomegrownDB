package buffer

import (
	"HomegrownDB/dbsystem/reldef"
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
	RTablePage(table reldef.TableRDefinition, pageId page.Id) (page.RPage, error)
	WTablePage(table reldef.TableRDefinition, pageId page.Id) (page.WPage, error)
}

type FsmBuffer interface {
	RFsmPage(ownerID reldef.OID, pageId page.Id) (fsmpage.Page, error)
	WFsmPage(ownerID reldef.OID, pageId page.Id) (fsmpage.Page, error)
}

const NewPage page.Id = page.InvalidId

type TableSrc interface {
	Table(id reldef.OID) reldef.TableRDefinition
}

type slotIndex = uint

var pageSize = int64(page.Size)
