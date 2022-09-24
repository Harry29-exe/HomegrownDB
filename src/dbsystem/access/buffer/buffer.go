package buffer

import (
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/schema/table"
)

var SharedBuffer DBSharedBuffer

func init() {
	SharedBuffer = NewSharedBuffer(10_000, table.DBTables)
}

//todo change methods to operate on ArrayIndexes
type DBSharedBuffer interface {
	RPage(tag bdata.PageTag) (bdata.RPage, error)
	WPage(id bdata.PageTag) (bdata.WPage, error)

	ReleaseWPage(tag bdata.PageTag)
	ReleaseRPage(tag bdata.PageTag)
}

type TableSrc interface {
	Table(id table.Id) table.Definition
}

type PageIO interface {
	Read(tag bdata.PageTag, buffer []byte)
	Flush(tag bdata.PageTag, buffer []byte)
	SaveNew(tag bdata.PageTag, buffer []byte)
}

type ArrayIndex = uint
