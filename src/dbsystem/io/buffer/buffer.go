package buffer

import (
	"HomegrownDB/dbsystem/bstructs"
	"HomegrownDB/dbsystem/io"
	"HomegrownDB/dbsystem/schema"
	"HomegrownDB/dbsystem/schema/table"
)

var SharedBuffer DBSharedBuffer = NewSharedBuffer(10_000, schema.Tables, io.Pages)

type DBSharedBuffer interface {
	RPage(tag bstructs.PageTag) (bstructs.RPage, error)
	WPage(id bstructs.PageTag) (bstructs.WPage, error)

	ReleaseWPage(tag bstructs.PageTag)
	ReleaseRPage(tag bstructs.PageTag)
}

type TableSrc interface {
	Table(id table.Id) table.Definition
}

type PageIO interface {
	Read(tag bstructs.PageTag, buffer []byte)
	Flush(tag bstructs.PageTag, buffer []byte)
	SaveNew(tag bstructs.PageTag, buffer []byte)
}

const bufferSize = 10_000

type ArrayIndex = uint
