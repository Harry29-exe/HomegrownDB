package buffer

import (
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/fsm/fsmpage"
	"HomegrownDB/dbsystem/storage/page"
)

type SharedBuffer interface {
	TableBuffer
	FsmBuffer
	SequenceBuffer

	WPageRelease(tag page.PageTag)
	RPageRelease(tag page.PageTag)
	FlushAll() error
}

type TableBuffer interface {
	RTablePage(table reldef.TableRDefinition, pageId page.Id) (page.TableRPage, error)
	WTablePage(table reldef.TableRDefinition, pageId page.Id) (page.TablePage, error)
}

type FsmBuffer interface {
	RFsmPage(ownerID reldef.OID, pageId page.Id) (fsmpage.Page, error)
	WFsmPage(ownerID reldef.OID, pageId page.Id) (fsmpage.Page, error)
}

type SequenceBuffer interface {
	SeqPage(seqOID reldef.OID) (page.SequencePage, error)
}

const NewPage page.Id = page.InvalidId
