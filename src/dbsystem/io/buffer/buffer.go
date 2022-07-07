package buffer

import (
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/io"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/stores"
)

var SharedBuffer DBSharedBuffer = NewSharedBuffer(10_000, stores.Tables, io.Pages)

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

const bufferSize = 10_000

type ArrayIndex = uint
