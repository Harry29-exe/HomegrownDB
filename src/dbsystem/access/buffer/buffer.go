package buffer

import (
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
	RPage(tag page.Tag) (page.TableRPage, error)
	WPage(id page.Tag) (page.TableWPage, error)

	ReleaseWPage(tag page.Tag)
	ReleaseRPage(tag page.Tag)
}

type TableSrc interface {
	Table(id table.Id) table.Definition
}

type ArrayIndex = uint
