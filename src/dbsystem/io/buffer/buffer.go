package buffer

import (
	"HomegrownDB/dbsystem/bstructs"
)

var Buffer DBBuffer = newBuffer(10_000)

type DBBuffer interface {
	RPage(tag bstructs.PageTag) (bstructs.RPage, error)
	WPage(id bstructs.PageTag) (bstructs.WPage, error)

	ReleaseWPage(tag bstructs.PageTag)
	ReleaseRPage(tag bstructs.PageTag)
}

const bufferSize = 10_000

type ArrayIndex = uint
