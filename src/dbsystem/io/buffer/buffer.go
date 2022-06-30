package buffer

import (
	"HomegrownDB/dbsystem/bstructs"
)

var Buffer DBBuffer = newBuffer(10_000)

type DBBuffer interface {
	RPage(tag bstructs.PageTag) (bstructs.RPage, error)
	WPage(id bstructs.PageTag) (bstructs.WPage, error)

	ReleaseWPage(page bstructs.WPage)
	ReleaseRPage(page bstructs.RPage)
}

const bufferSize = 10_000

type ArrayIndex = uint
