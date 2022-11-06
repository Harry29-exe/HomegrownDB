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
	RTablePage(table table.Definition, pageId page.Id) (tpage.TableRPage, error)
	WTablePage(table table.Definition, pageId page.Id) (tpage.TableWPage, error)
}

type FsmBuffer interface {
	RFsmPage(rel relation.Relation, pageId page.Id) (fsmpage.Page, error)
	WFsmPage(rel relation.Relation, pageId page.Id) (fsmpage.Page, error)
}

type TableSrc interface {
	Table(id table.Id) table.Definition
}

type arrayIndex = uint

func NewTablePageTag(pageIndex page.Id, tableDef table.Definition) page.Tag {
	return page.Tag{
		PageId:   pageIndex,
		Relation: tableDef.RelationID(),
	}
}

func NewPageTag(pageIndex page.Id, rel relation.Relation) page.Tag {
	return page.Tag{
		PageId:   pageIndex,
		Relation: rel.RelationID(),
	}
}
