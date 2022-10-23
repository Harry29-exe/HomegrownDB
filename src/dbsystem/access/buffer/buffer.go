package buffer

import (
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/access/dbbs"
	"HomegrownDB/dbsystem/schema/table"
)

var SharedBuffer DBSharedBuffer

func init() {
	SharedBuffer = NewSharedBuffer(10_000, table.DBTableStore, access.DBTableIOStore)
}

// todo change methods to operate on ArrayIndexes
type DBSharedBuffer interface {
	RPage(tag dbbs.PageTag) (dbbs.RPage, error)
	WPage(id dbbs.PageTag) (dbbs.WPage, error)

	ReleaseWPage(tag dbbs.PageTag)
	ReleaseRPage(tag dbbs.PageTag)
}

type TableSrc interface {
	Table(id table.Id) table.Definition
}

type PageIO interface {
	Read(tag dbbs.PageTag, buffer []byte)
	Flush(tag dbbs.PageTag, buffer []byte)
	SaveNew(tag dbbs.PageTag, buffer []byte)
}

type ArrayIndex = uint
