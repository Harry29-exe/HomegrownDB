package buffer

import (
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
)

var DBSharedBuffer SharedBuffer

func init() {
	DBSharedBuffer = NewSharedBuffer(10_000, pageio.DBStore)
}

// todo change methods to operate on ArrayIndexes
type SharedBuffer interface {
	RPage(tag PageTag) (Page, error)
	WPage(tag PageTag) (Page, error)

	ReleaseWPage(tag PageTag)
	ReleaseRPage(tag PageTag)
}

type TableSrc interface {
	Table(id table.Id) table.Definition
}

type ArrayIndex = uint

type Page = []byte

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
