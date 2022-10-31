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

	ReleaseWPage(tag PageTag)
	ReleaseRPage(tag PageTag)
}

type TableBuffer interface {
	TableRPage(tag PageTag, table table.Definition) (tpage.TableRPage, error)
	TableWPage(tag PageTag, table table.Definition) (tpage.TableWPage, error)
}

type FsmBuffer interface {
	RFsmPage(tag PageTag) (fsmpage.Page, error)
	WFsmPage(tag PageTag) (fsmpage.Page, error)
}

// todo change methods to operate on ArrayIndexes
type sharedBuffer interface {
	RPage(tag PageTag) (buffPage, error)
	WPage(tag PageTag) (buffPage, error)

	ReleaseWPage(tag PageTag)
	ReleaseRPage(tag PageTag)
}

type TableSrc interface {
	Table(id table.Id) table.Definition
}

type arrayIndex = uint

type buffPage struct {
	bytes []byte
	isNew bool
}

type PageTag struct {
	PageId   page.Id
	Relation relation.ID
}

func NewTablePageTag(pageIndex page.Id, tableDef table.Definition) PageTag {
	return PageTag{
		PageId:   pageIndex,
		Relation: tableDef.RelationId(),
	}
}

func NewPageTag(pageIndex page.Id, rel relation.Relation) PageTag {
	return PageTag{
		PageId:   pageIndex,
		Relation: rel.RelationID(),
	}
}
