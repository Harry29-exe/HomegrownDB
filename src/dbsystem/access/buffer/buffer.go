package buffer

import (
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/fsm/fsmpage"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/storage/tpage"
)

var DBSharedBuffer SharedBuffer

func init() {
	DBSharedBuffer = &bufferProxy{
		buffer: newSharedBuffer(10_000, pageio.DBStore),
	}
}

type SharedBuffer interface {
	TableBuffer
	FsmBuffer

	WPageRelease(tag page.Tag)
	RPageRelease(tag page.Tag)
}

type TableBuffer interface {
	RTablePage(pageId page.Id, table table.Definition) (tpage.TableRPage, error)
	WTablePage(pageId page.Id, table table.Definition) (tpage.TableWPage, error)
}

type FsmBuffer interface {
	RFsmPage(tag page.Tag) (fsmpage.Page, error)
	WFsmPage(tag page.Tag) (fsmpage.Page, error)
}

// todo change methods to operate on ArrayIndexes
type sharedBuffer interface {
	RPage(tag page.Tag) (buffPage, error)
	WPage(tag page.Tag) (buffPage, error)

	ReleaseWPage(tag page.Tag)
	ReleaseRPage(tag page.Tag)
}

type TableSrc interface {
	Table(id table.Id) table.Definition
}

type arrayIndex = uint

type buffPage struct {
	bytes []byte
	isNew bool
}

func NewTablePageTag(pageIndex page.Id, tableDef table.Definition) page.Tag {
	return page.Tag{
		PageId:   pageIndex,
		Relation: tableDef.RelationId(),
	}
}

func NewPageTag(pageIndex page.Id, rel relation.Relation) page.Tag {
	return page.Tag{
		PageId:   pageIndex,
		Relation: rel.RelationID(),
	}
}
